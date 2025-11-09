package main

import (
	"fmt"
	"gin-tut/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	str := fmt.Sprintf(":%v", os.Getenv("VALUE"))
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(str)

}
func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "H2llo!"})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}
}
