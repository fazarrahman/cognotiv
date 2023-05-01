package service

import (
	"context"
	"log"

	errorlib "github.com/fazarrahman/cognotiv/error"
	"github.com/fazarrahman/cognotiv/lib"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"golang.org/x/crypto/bcrypt"
)

// GetAccessTokenRequest ...
type GetAccessTokenRequest struct {
	Username string `validate:"required,min=1"`
	Password string `validate:"required,min=1"`
}

// UserPasswordCheckRequest ...
type UserPasswordCheckRequest struct {
	Username string `validate:"required,min=1"`
	Password string `validate:"required,min=1"`
}

// GetAccessToken ..
func (s *Svc) GetAccessToken(c *gin.Context, req *GetAccessTokenRequest) *errorlib.Error {
	res, err := s.CheckUsernamePassword(c, &UserPasswordCheckRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if res == nil && err != nil {
		return err
	}

	if *res == false {
		return errorlib.BadRequest("Invalid password")
	}

	clientId := lib.GetEnv("AUTH_CLIENT_ID")
	c.Request.ParseForm()
	c.Request.Form.Add("client_id", clientId)
	c.Request.Form.Add("client_secret", lib.GetEnv("AUTH_SECRET"))
	c.Request.Form.Add("scope", "read")
	c.Request.Form.Add("grant_type", "password")
	c.Request.Form.Add("username", req.Username)
	c.Request.Form.Add("password", req.Password)

	ginserver.SetPasswordAuthorizationHandler(
		func(ctx context.Context, clientId, username, password string) (userID string, err error) {
			us, errLib := s.UserRepository.GetUserByUsername(ctx, username)
			if err != nil {
				log.Println(errLib.Message)
				return "", nil
			}

			return us.Username, nil
		})

	ginserver.HandleTokenRequest(c)
	return err
}

// CheckUsernamePassword ..
func (s *Svc) CheckUsernamePassword(ctx *gin.Context, r *UserPasswordCheckRequest) (*bool, *errorlib.Error) {
	userEntity, err := s.UserRepository.GetUserByUsername(ctx, r.Username)
	var res bool
	if userEntity == nil && err == nil {
		return nil, errorlib.NotFound("No user data found")
	} else if err != nil {
		return nil, err
	}

	erro := bcrypt.CompareHashAndPassword(userEntity.Password, []byte(r.Password))
	if err != nil {
		res = false
		return &res, errorlib.InternalServerError(erro.Error())
	}
	res = true
	return &res, nil
}
