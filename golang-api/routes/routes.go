package routes

import (
    "golang-api/controllers"
    "golang-api/middlewares"

    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/login", controllers.Login)

    auth := r.Group("/")
    auth.Use(middlewares.JWTAuthMiddleware())
    {
        auth.POST("/terminal", controllers.CreateTerminal)
    }

    return r
}
