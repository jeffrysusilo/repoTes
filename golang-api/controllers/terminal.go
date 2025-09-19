package controllers

import (
    "golang-api/config"
    "golang-api/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

func CreateTerminal(c *gin.Context) {
    var input struct {
        Name string `json:"name"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    terminal := models.Terminal{Name: input.Name}
    if err := config.DB.Create(&terminal).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "terminal created", "terminal": terminal})
}
