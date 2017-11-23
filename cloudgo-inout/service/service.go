package service

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}

	mx.HandleFunc("/", homHandler(formatter)).Methods("GET")
	mx.HandleFunc("/regist", RegistHandler(formatter)).Methods("GET")
	mx.HandleFunc("/table", TableHandler(formatter)).Methods("POST")
	mx.HandleFunc("/unknown", UnknownHandler(formatter))
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(webRoot+"/assets/"))))
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
}

func homHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "23333", Content: "Hello Go!"})
	}
}

func RegistHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "regist", nil)
	}
}

func TableHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		formatter.HTML(w, http.StatusOK, "table", struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: req.Form["username"][0],
			Password: req.Form["password"][0],
		})
	}
}

func UnknownHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusNotImplemented, "501 Not Implemented")
	}
}
