package main

import (
    "net/http"
    "log"

    "github.com/jeffrysusilo/repotes/golang-api/config"
    "github.com/jeffrysusilo/repotes/golang-api/controllers"
    "github.com/jeffrysusilo/repotes/golang-api/middleware"
)

func main() {
    config.ConnectDB()

    http.HandleFunc("/login", controllers.Login)
    http.HandleFunc("/terminal", middleware.JwtVerify(controllers.CreateTerminal))

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
