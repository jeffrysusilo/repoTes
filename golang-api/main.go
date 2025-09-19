package main

import (
    "golang-api/config"
    "golang-api/models"
    "golang-api/routes"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    config.ConnectDB()
    config.DB.AutoMigrate(&models.User{}, &models.Terminal{})

    r := routes.SetupRouter()
    r.Run(":8080")
}
