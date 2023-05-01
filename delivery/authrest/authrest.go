package authRest

import (
	"log"
	"net/http"

	d "github.com/fazarrahman/cognotiv/delivery"
	"github.com/fazarrahman/cognotiv/model"
	"github.com/fazarrahman/cognotiv/service"

	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
)

// AuthRest ...
type AuthRest struct {
	Svc *service.Svc
}

// New ...
func New(_svc *service.Svc) *AuthRest {
	return &AuthRest{Svc: _svc}
}

// Register ...
func (r *AuthRest) Register(g *gin.RouterGroup) {
	g.POST("/login", r.Login)
	g.GET("/tokeninfo", d.MustAuthorize(), r.GetTokenInfo)
	g.POST("/signup", r.SignUp)
}

// GetTokenInfo ..
func (r *AuthRest) GetTokenInfo(c *gin.Context) {
	ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
	if exists {
		c.JSON(http.StatusOK, ti)
		return
	}
	log.Fatalln("error")
}

// Login ..
func (r *AuthRest) Login(c *gin.Context) {
	var req model.User
	c.BindJSON(&req)

	err := r.Svc.GetAccessToken(c, &service.GetAccessTokenRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	return
}

// SignUp ..
func (r *AuthRest) SignUp(c *gin.Context) {
	var req model.User
	c.BindJSON(&req)
	err := r.Svc.InsertUser(c, &req)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, "OK")
}
