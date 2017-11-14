package service

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	// 返回一个Render实例的指针
	formatter := render.New(render.Options{
		IndentJSON: true, // 输出时的格式是方便阅读的JSON
	})

	/* Classic 返回带有默认中间件的Negroni实例指针, 其中下面三项都实现了
	   Handler 接口的 ServeHTTP 方法:

	   Recovery - Panic Recovery Middleware
	   Logger - Request/Response Logging
	   Static - Static File Serving

	   func Classic() *Negroni {
	   	  return New(NewRecovery(), NewLogger(), NewStatic(http.Dir("public")))
	   }
	*/
	n := negroni.Classic()

	/* NewRouter 返回一个新的Router实例指针

	   func NewRouter() *Router {
		   return &Router{namedRoutes: make(map[string]*Route), KeepContext: false}
	   }

	   make(map[string]*Route) 为 namedRoutes 分配一个 map,
	   每个路径对应不同的处理函数(HandlerFunc)
	*/
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	// 让 negroni 使用该 Router
	// UseHandler adds a http.Handler onto the middleware stack. Handlers are invoked in the order they are added to a Negroni.
	/* func (n *Negroni) UseHandler(handler http.Handler) {
		n.Use(Wrap(handler))
	}
	*/
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	// 如果用户访问了地址 /hello/{id}, 那就对应调用该函数
	// 此处的函数为 testHandler 函数返回的 http.HandlerFunc

	// HandleFunc registers a new route with a matcher for the URL path.
	// See Route.Path() and Route.HandlerFunc().
	/*
		func (r *Router) HandleFunc(path string, f func(http.ResponseWriter,
			*http.Request)) *Route {
			return r.NewRoute().Path(path).HandlerFunc(f)
		}
	*/
	mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}
