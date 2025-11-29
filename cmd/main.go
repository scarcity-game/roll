package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/gaussian"
	"github.com/scarcity-game/roll/internal/uniform"
	"github.com/scarcity-game/roll/internal/weighted"
	"github.com/scarcity-game/roll/web"
)

func main() {
	router := gin.Default()
	router.GET("/rollDice", web.RollDice)
	router.GET("/sampleGaussian", gaussian.SampleGaussian)
	router.GET("/sampleUniform", uniform.SampleUniform)
	router.POST("/weightedChoice", weighted.MakeWeightedChoice)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("Unable to start server")
		fmt.Println(err)
		return
	}
}
