package main

import (
	"net/http"

	"example.com/event-mgmt/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.Run(":8090") // localhost:8090
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.BindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event Created Successfully!", "event": event})
}
