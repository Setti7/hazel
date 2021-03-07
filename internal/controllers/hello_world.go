package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HelloWorld(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"hello": "world!"})
}
