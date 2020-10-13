package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, c.Query("wechat"))
		//对于key对应value为空的 可以给一个默认值DefaultQuery
		c.String(200, c.DefaultQuery("id", "0"))
		//GetQuery和Query唯一不同的是，它返回两个值，可以告诉我们要获取的key是否存在。
		value, ok := c.GetQuery("wechat")
		if ok {
			c.String(200, value)
		}
	})
	r.Run(":8080")

}
