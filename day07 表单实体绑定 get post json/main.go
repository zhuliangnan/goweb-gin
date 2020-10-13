package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

//以一个用户注册功能来进行讲解表单实体绑定操作
//通过tag标签的方式设置每个字段对应的form表单中的属性名，通过binding属于设置属性是否是必须。
type UserRegister struct {
	Username string `form:"username"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

//使用ShouldBindQuery可以实现 Get 方式的数据请求的绑定。具体实现如下：
//使用ShouldBind可以实现Post方式的提交数据的绑定工作。具体编程如下所示：
//当客户端使用Json格式进行数据提交时，可以采用ShouldBindJson对数据进行绑定并自动解析,
// http://localhost:8080/hello?username=davie&&password=123&phone=19851782917
func main() {
	r := gin.Default()

	r.POST("/hello", func(c *gin.Context) {
		var user UserRegister
		//err := c.ShouldBindQuery(&user)
		err := c.ShouldBindJSON(&user)
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println(user.Username)
		fmt.Println(user.Password)
		fmt.Println(user.Phone)

		c.JSON(200, "成功")
	})

	r.Run(":8080")
}
