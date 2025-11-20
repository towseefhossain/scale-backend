package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    "scale-backend/handlers"
)

func main() {
    r := gin.Default()

    // CORS middleware
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // or specify domains like []string{"http://localhost:3000"}
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    PORT := ":8080"

    // API routes
    api := r.Group("/api/v1")
    {
        api.GET("/items", handlers.GetItems)
        api.POST("/items", handlers.CreateItem)
        api.GET("/items/:id", handlers.GetItem)
        api.PUT("/items/:id", handlers.UpdateItem)
        api.DELETE("/items/:id", handlers.DeleteItem)
    }

    log.Printf("Server starting on port %s...", PORT)
    r.Run(PORT)
}