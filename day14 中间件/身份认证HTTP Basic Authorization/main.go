package main

import "github.com/gin-gonic/gin"

func main() {

	//在Gin中，为我们提供了gin.BasicAuth帮我们生成基本认证的中间件，方便我们的开发。

	r := gin.Default()

	/*	r.Use(gin.BasicAuth(gin.Accounts{
			"admin":"123456",
		}))

		r.GET("/user", func(c *gin.Context) {
			c.JSON(200,"首页")
		})
	*/
	//分组路由来对于特定的url进行认证
	adminGroup := r.Group("/admin")

	adminGroup.Use(gin.BasicAuth(gin.Accounts{
		"admin": "123456",
	}))

	adminGroup.GET("/index", func(c *gin.Context) {
		c.JSON(200, "后台首页")
	})

	//针对特定URL的Basic Authorization

	r.Run(":8080")
}
