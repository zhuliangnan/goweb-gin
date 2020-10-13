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

		//请求处理 如果我们想输出业务处理结果的信息，该如何实现呢。答案是使用context.Next函数。
		//context.Next函数可以将中间件代码的执行顺序一分为二，Next函数调用之前的代码在请求处理之前之前，
		//当程序执行到context.Next时，会中断向下执行，转而先去执行具体的业务逻辑，执行完业务逻辑处理函数之后，
		//程序会再次回到context.Next处，继续执行中间件后续的代码。具体用法如下：
		c.Next()

		//处理后获取消耗时间  所以我们可以获取到程序执行完毕的时间
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("the request URL %s cost %v\n", url, costTime)
	}

}
