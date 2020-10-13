package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//统计每次请求的执行时间为例，来演示如何自定义一个中间件。

func main() {
	r := gin.New()

	r.Use(costTime())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "首页")
	})

	r.Run(":8080")
}

func costTime() gin.HandlerFunc {

	return func(c *gin.Context) {
		//请求前获取当前时间
		nowTime := time.Now()

		//请求处理
		c.Next()

		//处理后获取消耗时间
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}

}
