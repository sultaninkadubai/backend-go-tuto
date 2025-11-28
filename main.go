package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
    	c.JSON(200, gin.H{
			"hello": "world",
		})																			
    })
    if err := r.Run(":9090"); err !=nil {
		log.Fatalf("failed to run server: %v", err)
	}																															
}
