package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/web"
)

func main() {
	router := gin.Default()
	router.GET("/rollDice", web.RollDice)
	router.GET("/sampleGaussian", web.SampleGaussian)
	router.GET("/sampleUniform", web.SampleUniform)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Unable to start server")
		fmt.Println(err)
		return
	}
}
