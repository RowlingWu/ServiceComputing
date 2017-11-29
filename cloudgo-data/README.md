使用`xorm`框架

# 实现效果

## mysql中表的结构

![userinfo 表的结构](http://img.blog.csdn.net/20171129223125609?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

## 插入数据
插入数据内容为：`username="Service", departname="sdcs"`

![插入数据](http://img.blog.csdn.net/20171129223526199?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

## 访问表中所有数据

![查询数据库1](http://img.blog.csdn.net/20171129225834965?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


![查询数据库2](http://img.blog.csdn.net/20171129223926143?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

## 查询一项条目

![查询一项条目](http://img.blog.csdn.net/20171129224128226?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

# 对比 xorm 和 database/sql

## orm 是否实现了 dao 的自动化？

我认为`xorm`应该是实现了dao的自动化。通过使用 xorm, 程序员省去了实现 dao 接口的步骤，同时每次使用一个 Query, xorm 都已帮我们写好了 Close 函数，程序员不需要使用 `defer Stmt.Close()`.


## ab 性能测试

### xorm 性能测试

![xorm测试1](http://img.blog.csdn.net/20171129225021120?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![xorm测试2](http://img.blog.csdn.net/20171129225044438?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

### database/sql 性能测试

![sql 测试1](http://img.blog.csdn.net/20171129225357578?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![sql 测试2](http://img.blog.csdn.net/20171129225412834?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

**结论：** database/sql 的响应速度要比 xorm 快