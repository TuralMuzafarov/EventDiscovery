package routes

import (
	"github.com/TuralMuzafarov/EventDiscovery/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	authenticated := router.Group("/")
	authenticated.Use(middlewares.Authenticate)
	{
		authenticated.POST("/events", createEvent)
		authenticated.PUT("/events/:id", updateEvent)
		authenticated.DELETE("/events/:id", deleteEvent)
	}
	router.GET("/events", getEvents)
	router.GET("events/:id", getEvent) // events/1, /events/5
	router.POST("/signup", signUp)
	router.POST("/login", login)
}
