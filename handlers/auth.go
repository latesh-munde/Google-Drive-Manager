package handlers

import (
	"context"
	"encoding/json"
	"os"

	"gin-drive-manager/config"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func Login(c *gin.Context) {
	conf := config.GetOAuthConfig()

	url := conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(302, url)
}

func Callback(c *gin.Context) {
	conf := config.GetOAuthConfig()
	code := c.Query("code")

	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(500, gin.H{"error": "Token exchange failed"})
		return
	}

	file, _ := os.Create("tokens/token.json")
	defer file.Close()
	json.NewEncoder(file).Encode(token)

	c.JSON(200, gin.H{"message": "Authentication successful"})
}
