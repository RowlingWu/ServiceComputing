package main

import (
	"github.com/RowlingWu/cloudgo/service"
)

func main() {
	// NewServer func() *negroni.Negroni
	// NewServer 返回一个Negroni实例的指针.
	server := service.NewServer()

	// Run 最后会调用 http.ListenAndServe, 并把 negroni 作为 Handler 接口传入
	// 于是 negroni 和 negroni.middleware 实现的 ServeHTTP 方法会被先后调用
	//
	// func (n *Negroni) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	//    n.middleware.ServeHTTP(NewResponseWriter(rw), r)
	// }
	//
	// func (m middleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// 开始调用链表中的第一个 handler 的 ServeHTTP 方法
	//  	m.handler.ServeHTTP(rw, r, m.next.ServeHTTP)
	// }
	server.Run(":8000")
}
