package main

import (
    "net/http"
    "log"

    "golang-api/config"
    "golang-api/controllers"
    "golang-api/middleware"

)

func main() {
    config.ConnectDB()

    http.HandleFunc("/login", controllers.Login)
    http.HandleFunc("/terminal", middleware.JwtVerify(controllers.CreateTerminal))

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
