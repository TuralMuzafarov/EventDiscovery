package main

import (
	"net/http"

	"github.com/TuralMuzafarov/EventDiscovery/db"
	"github.com/TuralMuzafarov/EventDiscovery/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not retrieve events",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	event.UserID = 1
	event.ID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}
