package middleware

import (
	"gabefraser/minepass/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorcon/rcon"
)

func RconMiddleware(c *gin.Context) {
	host := os.Getenv("MP_HOST")
	if host == "" {
		utils.Logger("Missing MP_HOST environment variable, exiting...")
		os.Exit(1)
	}

	password := os.Getenv("MP_PASSWORD")
	if password == "" {
		utils.Logger("Missing MP_PASSWORD environment variable, exiting...")
		os.Exit(1)
	}

	conn, err := rcon.Dial(host, password)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c.Set("rcon", conn)
	c.Next()
}
