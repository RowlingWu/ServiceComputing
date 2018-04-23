package service

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"session"
	_ "session/memory" // 该包给session的manager注册一个自实现的provider，名为memory
	"strconv"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// globalSessons 全局的Session管理器
var globalSessions *session.Manager

func init() {
	globalSessions, _ = session.NewManager("memory", "goSessionID", session.MAX_LIFE_TIME)
	go globalSessions.GC()
}

var prefix = "html/"

func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx)

	n.UseHandler(mx)
	n.UseFunc(printManager())
	return n
}

func initRoutes(mx *mux.Router) {
	mx.HandleFunc("/login", login)
	mx.HandleFunc("/upload", upload)
	mx.HandleFunc("/", homeHandler)
	mx.HandleFunc("/count", count)
	mx.HandleFunc("/logout", logout)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	t, _ := template.ParseFiles(prefix + "index.html")
	t.Execute(w, sess.Get("username"))
}

func printManager() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		globalSessions.Print()
		next(w, r)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	globalSessions.SessionDestroy(w, r)
	http.Redirect(w, r, "/", 302)
}

func login(w http.ResponseWriter, req *http.Request) {
	sess := globalSessions.SessionStart(w, req)
	req.ParseForm()
	if req.Method == "GET" {
		t, _ := template.ParseFiles(prefix + "login.html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", req.Form["username"])
		http.Redirect(w, req, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles(prefix + "count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

func upload(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles(prefix + "upload.gtpl")
		t.Execute(w, token)
	} else {
		req.ParseMultipartForm(32 << 20) // 上传的文件存储在maxMemory大小的内存里面
		// 如果文件大小超过了maxMemory，那么剩下的部分将存储在系统的临时文件中
		file, handler, err := req.FormFile("uploadfile") // 获取上面的文件句柄
		if err != nil {
			panic(err)
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./uploadFiles/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
