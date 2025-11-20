package handlers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

// GetItems handles GET requests to retrieve all items
func GetItems(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"items": [3]int{1, 2, 3}})
}

// CreateItem handles POST requests to create a new item
func CreateItem(c *gin.Context) {
    c.JSON(http.StatusCreated, "item created")
}

// GetItem handles GET requests for a specific item
func GetItem(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"item": id})
}

// UpdateItem handles PUT requests to update an item
func UpdateItem(c *gin.Context) {
    c.JSON(http.StatusOK, "item updated")
}

// DeleteItem handles DELETE requests to remove an item
func DeleteItem(c *gin.Context) {
    c.JSON(http.StatusOK, "item deleted")
}