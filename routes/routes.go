package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.GET("events/:id", getEvent) // events/1, /events/5
	router.POST("/events", createEvent)
	router.PUT("/events/:id", updateEvent)
	router.DELETE("/events/:id", deleteEvent)
}
