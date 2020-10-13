package main

import "github.com/gin-gonic/gin"

func main() {

	//托管一个静态文件
	//在项目的开发中，你可能需要这么一个功能：把服务器上的JS文件暴露出来以供访问，比如让网站调用里面的JS函数等。对于这种情形，我们可以使用Gin提供的StaticFile方法很方便的完成。

	//把文件/tmp/adobegc.log托管在网络上，并且设置访问路径为/adobegc.log,这样我们通过http://localhost:8080/adobegc.log就可以访问这个文件，看到它的内容了。
	router := gin.Default()
	router.StaticFile("/adobegc.log", "tmp/adobegc.log")

	//托管一个目录
	router.Static("/static", "tmp")

	router.Run(":8080")

}
