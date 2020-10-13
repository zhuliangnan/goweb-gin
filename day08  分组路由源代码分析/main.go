package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/** 分组路由
/admin/users
/admin/manager
/admin/photo
*/
func main() {
	r := gin.Default()

	//V1版本的API
	v1Group := r.Group("/v1", func(c *gin.Context) {
		fmt.Println("/v1中间件")
	})

	v1Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v1/users")
	})
	v1Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v1/products")
	})

	//V2版本的API
	v2Group := r.Group("/v2")
	v2Group.GET("/users", func(c *gin.Context) {
		c.String(200, "/v2/users")
	})
	v2Group.GET("/products", func(c *gin.Context) {
		c.String(200, "/v2/products")
	})

	//在v1Group这个分组路由中再添加一个分组路由
	v1AdminGroup := v1Group.Group("/admin")
	{
		v1AdminGroup.GET("/users", func(c *gin.Context) {
			c.String(200, "/v1/admin/users")
		})
		v1AdminGroup.GET("/manager", func(c *gin.Context) {
			c.String(200, "/v1/admin/manager")
		})
		v1AdminGroup.GET("/photo", func(c *gin.Context) {
			c.String(200, "/v1/admin/photo")
		})
	}

	r.Run(":8080")

}
