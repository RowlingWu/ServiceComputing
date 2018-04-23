package memory

import (
	"container/list"
	"fmt"
	"log"
	"session"
	"sync"
	"time"
)

var pder = &Provider{list: list.New()}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
}

// SessionStore 就是一个个session
type SessionStore struct {
	sid          string                      // session id 唯一标识
	timeAccessed time.Time                   // 最后访问时间
	value        map[interface{}]interface{} // session里面存储的值
}

// Set 给该session设值，map[key] = value，并更新该session的最后访问时间
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

// Get 返回该session的key所对应的value，并更新该session的最后访问时间
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	}
	return nil
}

// Delete 删除该session的某个key和value，更新该session的最后访问时间
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

// SessionID 返回该session的sid
func (st *SessionStore) SessionID() string {
	return st.sid
}

// Provider 管理所有的session。包含的操作有
// SessionInit 接受一个sid，创建并返回一个新的Session
// SessionRead 接受一个sid，若该sid的session存在，返回它；若不存在，调用SessionInit并返回一个新的session
// SessionDestroy 从map和list中删除该session
// SessionGC 删除所有过期的session
// SessionUpdate 更新sid对应的session的最后访问时间（更新为time.Now()）
type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element // map[sid]&SessionStore
	list     *list.List
}

func (pder *Provider) Print() {
	for element := pder.list.Front(); element != nil; element = element.Next() {
		log.Println("\n", element.Value.(*SessionStore))
	}
}

func (st *SessionStore) String() string {
	var str string
	str += fmt.Sprintln("sid:", st.sid)
	str += fmt.Sprintln("timeAccess:", st.timeAccessed)
	str += fmt.Sprintln("value(map):\n", st.value)
	return str
}

// SessionInit 接受一个sid，创建并返回一个新的Session
func (pder *Provider) SessionInit(sid string) (session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

// SessionRead 接受一个sid，若该sid的session存在，返回它；若不存在，调用SessionInit并返回一个新的session
func (pder *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
}

// SessionDestroy 从map和list中删除该session
func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

// SessionGC 删除所有过期的session
func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

// SessionUpdate 更新sid对应的session的最后访问时间（更新为time.Now()）
func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}
