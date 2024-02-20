package handlers

import (
	"gabefraser/minepass/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorcon/rcon"
)

func WhitelistAdd(c *gin.Context) {
	rcon := c.MustGet("rcon").(*rcon.Conn)

	var req WhitelistUserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No username was provided. Please try again.",
		})
		return
	}

	response, err := rcon.Execute("whitelist add " + req.Username)
	if err != nil {
		log.Fatal(err)
	}

	utils.Logger(c.ClientIP() + " added " + req.Username + " to the whitelist")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": response,
	})
}

func WhitelistRemove(c *gin.Context) {
	rcon := c.MustGet("rcon").(*rcon.Conn)

	var req WhitelistUserRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger("No username was provided from " + c.ClientIP())

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "No username was provided. Please try again.",
		})
		return
	}

	response, err := rcon.Execute("whitelist remove " + req.Username)
	if err != nil {
		log.Fatal(err)
	}

	utils.Logger(c.ClientIP() + " removed " + req.Username + " from the whitelist")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": response,
	})
}
