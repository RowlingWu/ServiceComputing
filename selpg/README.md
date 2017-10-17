#selpg#
a command line 

##Usage##
`selpg  -s start_page -e end_page [ -f | -l lines_per_page ] [ -d dest ] [ in_filename ]`

##Descriptions of options##
`-s` and `-e` must be needed

`-d` should be followed by a file name whose content is the output result. If `-d` is followed by nothing, `Stdout` is the default output.

`in_filename` is the input file name, and if you don't input it, the default input will be `Stdin`.

#Outputs#
The file `input` can be seen in the directory `/selpg` 

 1. `$ selpg -s 1 -e 1 < input`
![这里写图片描述](http://img.blog.csdn.net/20171017234437002?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 2. `$ selpg -s 1 -e 1 input`
![这里写图片描述](http://img.blog.csdn.net/20171017234609164?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 3. `$ cat input | selpg -s 1 -e 2`
![这里写图片描述](http://img.blog.csdn.net/20171017234816880?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 4. `$ selpg -s 1 -e 1 input > output`
![这里写图片描述](http://img.blog.csdn.net/20171017234953736?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171017235041445?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 5. `$ selpg -s 10 -e 20 input > 2>errfile`
 There is no error so the file `errfile` is created but empty.
 There is nothing to be printed out so the `Stdout` is empty.
![这里写图片描述](http://img.blog.csdn.net/20171017235538544?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 6. `$ selpg -s 10 -e 20 input >output 2>errfile`
  There is no error so the file `errfile` is created but empty.
![这里写图片描述](http://img.blog.csdn.net/20171018000007116?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171018000126183?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 7. `$ selpg -s 1 -e 1 input | cat`
![这里写图片描述](http://img.blog.csdn.net/20171018000525332?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 8. `$ selpg -s 1 -e 3 -l 2 input`
![这里写图片描述](http://img.blog.csdn.net/20171018000742944?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 9. `$ selpg -s 1 -e 2 -f input`
![这里写图片描述](http://img.blog.csdn.net/20171018000939966?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 10. `$ selpg -s 1 -e 1 -d output input`
![这里写图片描述](http://img.blog.csdn.net/20171018001109146?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
![这里写图片描述](http://img.blog.csdn.net/20171018001142411?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 11. `$ selpg -s 1 -e 2 -f -l 10 input 2>errfile`
![这里写图片描述](http://img.blog.csdn.net/20171018002447262?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171018002607695?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

 12.`$ selpg -s 2 -e 3 -l 2 `
 ![这里写图片描述](http://img.blog.csdn.net/20171018003046647?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
 13. `$ selpg -s 2 -e 3 -l 2 -d output`
![这里写图片描述](http://img.blog.csdn.net/20171018003317703?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171018003402176?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvd3VybGlu/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)