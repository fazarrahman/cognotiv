package lib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
)

// GetEnv ...
func GetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Printf("Env %s value not exist \n", key)
	return ""
}

// GetUsernameFromToken ...
func GetUsernameFromToken(c *gin.Context) *string {
	ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
	if exists {
		a := strings.Split(fmt.Sprintf("%v", ti), " ")
		return &a[1]
	}

	return nil
}
