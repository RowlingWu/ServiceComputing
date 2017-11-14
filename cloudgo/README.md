# 说明
代码中的注释较为粗略，更为详细的解读在
http://blog.csdn.net/wurlin/article/details/78535983

# 使用框架
使用 `gorilla/mux` 和 `negroni`

## 决策依据
negroni 适用于一般 web 应用与服务开发，而在 negroni 的 README.md 中提到：

> Negroni 没有带路由功能，使用 Negroni 时，需要找一个适合你的路由。不过好在 Go 社区里已经有相当多可用的路由，Negroni 更喜欢和那些完全支持 net/http 库的路由搭配使用，比如搭配 Gorilla Mux 路由器

既然是官方推荐，自有其原因，开发时遇到的问题相对来说应该也少。再者， gorilla/mux 包可以方便程序员使用正则表达式、提取 URL path 中参数，比原生的 http/net 包使用方便。

# Curl 测试
![curl 测试](http://img.blog.csdn.net/20171114230116278?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

# ab 测试
## 参数解释
-n 发出的测试请求数目。截图中为1000个。
-c 每次并发的测试请求数目。截图中为100个。

![ab测试](http://img.blog.csdn.net/20171114230459938?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


![ab测试续](http://img.blog.csdn.net/20171114230446619?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)