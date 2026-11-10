package main

import (
	"github.com/TuralMuzafarov/EventDiscovery/db"
	"github.com/TuralMuzafarov/EventDiscovery/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
