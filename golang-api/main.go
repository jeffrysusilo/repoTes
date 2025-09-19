package main

import (
    "net/http"
    "log"

    "github.com/yourusername/go-ticket/config"
    "github.com/yourusername/go-ticket/controllers"
    "github.com/yourusername/go-ticket/middleware"
)

func main() {
    config.ConnectDB()

    http.HandleFunc("/login", controllers.Login)
    http.HandleFunc("/terminal", middleware.JwtVerify(controllers.CreateTerminal))

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
