package delivery

import (
	"log"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
)

// MustAuthorize ..
func MustAuthorize() gin.HandlerFunc {
	return ginserver.HandleTokenVerify(ginserver.Config{ErrorHandleFunc: ginserver.ErrorHandleFunc(func(c *gin.Context, e error) {
		log.Println(e.Error())
		c.AbortWithStatusJSON(401, "Not Authorized: "+e.Error())
		return
	})})

}
