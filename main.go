package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "scale-backend/handlers"
)

func main() {
    r := gin.Default()

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

    log.Println("Server starting on port %s...", PORT)
    r.Run(PORT)
}