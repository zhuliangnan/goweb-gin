package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	router := gin.Default()

	router.GET("/adobegc.log", func(c *gin.Context) {
		data, err := ioutil.ReadFile("tmp/adobegc.log")
		if err != nil {
			c.AbortWithError(500, err)
		} else {
			//可以指定contentType和内容data   这种能力很有用，比如我们可以把我们储存在数据库中的图片二进制数据，作为一张图片显示在网站上。
			c.Data(200, "text/plain; charset=utf-8", data)
		}
	})

	router.Run(":8080")
}
