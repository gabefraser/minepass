package handlers

import "github.com/gin-gonic/gin"

func Validate(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "Password is valid",
	})
}
