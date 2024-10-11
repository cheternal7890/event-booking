package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	
	server.POST("/events", createEvent) // Auth Required
	server.PUT("/events/:id", updateEvent) // Auth Required
	server.DELETE("/events/:id", deleteEvent) // Auth Required

	server.POST("/signup", signup)
	server.POST("/login", login)
}
