package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/internal/dice"
	"github.com/scarcity-game/roll/internal/gaussian"
	"github.com/scarcity-game/roll/internal/uniform"
	"github.com/scarcity-game/roll/internal/weighted"
)

func main() {
	router := gin.Default()
	router.GET("/rollDice", dice.RollDice)
	router.GET("/sampleGaussian", gaussian.SampleGaussian)
	router.GET("/sampleUniform", uniform.SampleUniform)
	router.POST("/weightedChoice", weighted.MakeWeightedChoice)

	err := router.Run("0.0.0.0:8080")
	if err != nil {
		fmt.Println("Unable to start server")
		fmt.Println(err)
		return
	}
}
