package main

import (
    "github.com/Setti7/hazel/internal/controllers"
    "github.com/Setti7/hazel/internal/middlewares"
    "github.com/Setti7/hazel/internal/models"
    "github.com/Setti7/hazel/internal/setup"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    setup.ConnectDatabase()
    setup.SetupIdTokenVerifier()

    r.Use(gin.Logger())
    r.Use(middlewares.Authorize())

    r.GET("/healthz", func(c *gin.Context) {
        user := c.MustGet("user").(*models.User)
        c.JSON(http.StatusOK, gin.H{"data": user})
    })

    r.GET("/", controllers.HelloWorld)

    r.Run("0.0.0.0:8000")
}
