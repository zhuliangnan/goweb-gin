package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	//QueryArray  --  http://localhost:8080/?media=blog&media=wechat
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, c.QueryArray("media"))
	})

	//QueryMap ---  http://localhost:8080/map?ids[a]=123&ids[b]=456&ids[c]=789
	r.GET("/map", func(c *gin.Context) {
		c.JSON(200, c.QueryMap("ids"))
	})
	r.Run(":8080")
}
