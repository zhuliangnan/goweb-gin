package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()
	//自定义函数 	这里以自己实现一个md5加密为例，演示自定义模板函数的使用。
	//SetFuncMap方法把我们自定义的函数Set进去，并且要在LoadHTMLGlob加载模板之前，不然不起作用
	r.SetFuncMap(template.FuncMap{
		"md5": MD5,
	})
	//首先创建了一个index.html模板，放在html文件夹下,内容很简单。
	//html/index.html --- 微信公众号: {{printf "%s" .}}
	//然后通过r.LoadHTMLFiles("html/index.html")加载这个模板文件，这样我们才能使用它。
	//从根目录加载
	r.LoadHTMLFiles("html/index.html")
	//加载一个目录下的所有模板。
	//r.LoadHTMLFiles("html/*")
	//模板目录结构是这样的： html/user/index.html  html/admin/index.html html/product/index.html
	//r.LoadHTMLFiles("html/**/*")
	r.GET("/html", func(c *gin.Context) {
		//使用的方式非常简单，通过c.HTML(200, "index.html", "flysnow_org")即可。
		c.HTML(200, "index.html", "flysnow_org")
	})

	r.Run(":8080")
}

func MD5(in string) (string, error) {
	hash := md5.Sum([]byte(in))
	return hex.EncodeToString(hash[:]), nil
}
