package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/jsonp", func(c *gin.Context) {
		c.JSONP(200, gin.H{"wechat": "flysnow_org"})
	})
	//解决JSON劫持
	a := []string{"1", "2", "3"}
	//默认while 我们还可以手动更改策略
	r.SecureJsonPrefix("for(;;);")
	r.GET("/secureJson", func(c *gin.Context) {
		c.SecureJSON(200, a)
	})
	r.Run(":8080")
}
