package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{"wechat": "flysnow_org", "blog": "www.flysnow.org"})
	})
	r.GET("/xmlUser", func(c *gin.Context) {
		c.XML(200, User{ID: 123, Name: "张三", Age: 20})
	})
	//xml数组
	//因为XML必须要有一个根节点，所以我们必须要有一个对象存放我们的struct数组，比如map.
	r.GET("/xmlArray", func(c *gin.Context) {
		allUsers := []User{{ID: 123, Name: "张三", Age: 20}, {ID: 456, Name: "李四", Age: 25}}
		c.XML(200, gin.H{"user": allUsers})
	})

	r.Run(":8080")
}

type User struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	Age  int    `xml:"age"`
}
