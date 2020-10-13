package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := []User{
		{ID: 123, Name: "zhuzhu"},
		{ID: 345, Name: "zhu2"},
	}
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.POST("/users", func(context *gin.Context) {
		//创建一个用户
	})
	r.DELETE("/usrs/123", func(context *gin.Context) {
		//删除ID为123的用户
	})
	r.PUT("/usrs/123", func(context *gin.Context) {
		//更新ID为123的用户
	})

	r.PATCH("/usrs/123", func(context *gin.Context) {
		//更新ID为123用户的部分信息
	})
	r.Run(":8080")
}
