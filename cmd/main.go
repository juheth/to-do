package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juheth/to-do/core/user"
	"github.com/juheth/to-do/db"
)

func main() {

	db, err := db.ConnectionDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	userRepo := user.NewRepo(db)
	userSrv := user.NewService(userRepo)
	userEnd := user.MakeEnponints(userSrv)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/register", gin.HandlerFunc(userEnd.RegisterUser))
	router.GET("/users", gin.HandlerFunc(userEnd.GetUser))

	router.Run(":8080")

}
