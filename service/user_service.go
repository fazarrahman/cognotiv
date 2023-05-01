package service

import (
	"time"

	ue "github.com/fazarrahman/cognotiv/domain/user/entity"
	"github.com/fazarrahman/cognotiv/error"
	errorlib "github.com/fazarrahman/cognotiv/error"
	"github.com/fazarrahman/cognotiv/model"
	"github.com/ulule/deepcopier"

	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	validate "gopkg.in/go-playground/validator.v9"
)

func (s *Svc) InsertUser(ctx *gin.Context, r *model.User) *errorlib.Error {
	errs := validate.New().Struct(r)
	if errs != nil {
		log.Println(errs)
		return error.BadRequest(errs.Error())
	}

	pwdhash, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return error.InternalServerError(err.Error())
	}

	// validate username
	us, errL := s.UserRepository.GetUserByUsername(ctx, r.Username)
	if err != nil {
		log.Println(errL.Message)
		return errL
	}
	if us != nil {
		return error.ResourceAlreadyExist("Username already exists")
	}

	// validate email
	eEmail, errL := s.UserRepository.GetUserByEmail(ctx, r.Email)
	if err != nil {
		log.Println(errL.Message)
		return errL
	}
	if eEmail != nil {
		return error.ResourceAlreadyExist("Email is already registered")
	}

	var u ue.Users
	deepcopier.Copy(r).To(&u)
	u.Password = pwdhash
	u.CreatedDate = time.Now()
	errL = s.UserRepository.InsertUser(ctx, &u)

	if errL != nil {
		log.Println(errL)
		return errL
	}

	return nil
}
