package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()
	router.StaticFS("/static1", gin.Dir("tmp", true))
	router.Run(":8080")

}
