package main

import (
    "golang-api/config"
    "golang-api/models"
    "log"
    "time"

    "golang.org/x/crypto/bcrypt"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    config.ConnectDB()

    config.DB.AutoMigrate(&models.User{}, &models.Terminal{})

    password := "admin123"
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

    admin := models.User{
        Username: "admin",
        Password: string(hashedPassword),
        Model:    models.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()},
    }

    if err := config.DB.Create(&admin).Error; err != nil {
        log.Println("Admin user already exists or error:", err)
    } else {
        log.Println("Admin user created with username: admin and password:", password)
    }

    terminals := []models.Terminal{
        {Name: "Terminal A"},
        {Name: "Terminal B"},
        {Name: "Terminal C"},
    }

    for _, t := range terminals {
        if err := config.DB.Create(&t).Error; err != nil {
            log.Println("Terminal", t.Name, "already exists or error:", err)
        } else {
            log.Println("Terminal created:", t.Name)
        }
    }

    log.Println("Seeding completed")
}
