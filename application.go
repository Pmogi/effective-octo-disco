package main

import (
	"GoTweet/Controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	messageApi := Controller.MessageApi{Engine: r}
	messageApi.Init()

	err := r.Run(":5000")
	if err != nil {
		panic(err)
	}
}
