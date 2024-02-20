package middleware

import (
	"gabefraser/minepass/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func GuardMiddleware(c *gin.Context) {
	password := os.Getenv("MP_UI_PASSWORD")

	apiKey := c.GetHeader("X-Api-Key")
	if apiKey == "" {
		utils.Logger("Missing password from the request from " + c.ClientIP())

		c.JSON(401, gin.H{
			"success": false,
			"message": "Missing password from the request",
		})
		c.Abort()
		return
	}

	if apiKey != password {
		utils.Logger("Incorrect password from " + c.ClientIP())

		c.JSON(401, gin.H{
			"success": false,
			"message": "Incorrect password",
		})
		c.Abort()
		return
	}

	c.Next()
}
