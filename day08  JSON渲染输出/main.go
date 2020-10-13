package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})

	//可以把我们自定义的对象struct转为一个json字符串输出
	r.GET("/users/123", func(c *gin.Context) {
		c.JSON(200, user{ID: 123, Name: "张三", Age: 20})
	})
	//json数组输出需要传输给他一个数组就可以了  IndentedJSON---json美化
	allUsers := []user{{ID: 123, Name: "张三", Age: 20}, {ID: 234, Name: "里斯", Age: 21}}
	r.GET("/users", func(c *gin.Context) {
		c.IndentedJSON(200, allUsers)
	})

	//PureJSON -- 不转义
	/*对于JSON字符串中特殊的字符串，比如<，Gin默认是转义的，比如变成\ u003c，
	但是有时候我们为了可读性，需要保持原来的字符，不进行转义，这时候我们就可以使用PureJSON*/
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "<b>Hello, world!</b>",
		})
	})

	r.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"message": "<b>Hello, world!</b>",
		})
	})

	//AsciiJSON -- 转成unicode编码
	//如果要把非Ascii字符串转为unicode编码，Gin同样提供了非常方便的方法。
	r.GET("/asciiJSON", func(c *gin.Context) {
		c.AsciiJSON(200, gin.H{"message": "hello 飞雪无情"})
	})

	//加速JSON
	/*在Gin中，提供了两种JSON解析器，用于生成JSON字符串。默认的是Golang(Go语言)内置的JSON，
	当然你也可以使用jsoniter，据说速度很快。如果要使用jsoniter，我们在go build编译的时候只需要这么做即可：
	go build -tags=jsoniter .
	*/

	r.Run(":8080")
}

type user struct {
	//自定义JSON字段名称
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
