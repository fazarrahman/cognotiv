package main

import (
	"log"
	"net/http"

	"github.com/fazarrahman/cognotiv/auth"
	"github.com/fazarrahman/cognotiv/config/mysqldb"
	auth_rest "github.com/fazarrahman/cognotiv/delivery/authrest"
	delivery_rest "github.com/fazarrahman/cognotiv/delivery/rest"
	order_repo_db "github.com/fazarrahman/cognotiv/domain/order/repository/mysqldb"
	user_repo_db "github.com/fazarrahman/cognotiv/domain/user/repository/mysqldb"
	"github.com/fazarrahman/cognotiv/lib"
	"github.com/fazarrahman/cognotiv/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envInit()

	db, err := mysqldb.New()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database has been initialized")

	auth.Init()
	log.Println("Oauth has been initialized")

	orderDb := order_repo_db.New(db)
	userDb := user_repo_db.New(db)

	svc := service.New(orderDb, userDb)

	g := gin.Default()
	g.GET("/ping",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	delivery_rest.New(svc).Register(g.Group("/api"))
	auth_rest.New(svc).Register(g.Group("/api/auth"))

	g.Run(":" + lib.GetEnv("APP_PORT"))
}

func envInit() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatalln("No .env file found")
	}
}
