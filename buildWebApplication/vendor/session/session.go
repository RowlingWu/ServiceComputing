package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Manager 是全局的管理器
type Manager struct {
	cookieName  string
	lock        sync.Mutex // protects session
	provider    Provider
	maxLifeTime int64
}

const MAX_LIFE_TIME int64 = 60 * 1

func (manager *Manager) Print() {
	manager.provider.Print()
}

// sessionID 返回全局唯一的sessionID
func (manager *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionStart 检测是否已经有某个Session与当前来访用户发生了关联，如果没有则创建之
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
		cookie.Expires = time.Time.Add(time.Now(), time.Second*time.Duration(MAX_LIFE_TIME))
		http.SetCookie(w, cookie)
	}
	return
}

// SessionDestroy 删除manager.provider中对应的cookie，删除浏览器的cookie
func (manager *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	sid, _ := url.QueryUnescape(cookie.Value)
	manager.provider.SessionDestroy(sid)
	expiration := time.Now()
	cookie2 := http.Cookie{Name: manager.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
	http.SetCookie(w, &cookie2)
}

func (manager *Manager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	log.Println("sessionGC\n")
	time.AfterFunc(time.Duration(manager.maxLifeTime*(int64)(time.Second)), func() { manager.GC() }) // 利用time包的定时器功能
}

// NewManager 返回一个全局的管理器Manager
func NewManager(provideName, cookieName string, maxLifeTime int64) (*Manager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

var provides = make(map[string]Provider)

// Register makes a session provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

type Provider interface {
	SessionInit(sid string) (Session, error) // 实现Session的初始化，返回新的Session变量
	SessionRead(sid string) (Session, error) // 返回sid对应的Session。若不存在该Session，调用SessionInit()
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64) // 根据maxLifeTime删除过期的数据
	Print()
}

type Session interface {
	Set(key, value interface{}) error // set session value
	Get(key interface{}) interface{}  // get session value
	Delete(key interface{}) error     // delete session value
	SessionID() string                // 获取当前sessionID
}
