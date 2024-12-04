package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/juheth/to-do/core/middleware"
	"github.com/juheth/to-do/core/user"
	"github.com/juheth/to-do/db"
)

func main() {

	db, err := db.ConnectionBD()

	if err != nil {
		log.Fatalf(err.Error())
	}

	userRepo := user.NewRepo(db)
	userSrv := user.NewService(userRepo)
	userEnd := user.MakeEnponints(userSrv)

	// taskRepo := task.NewRepo(db)
	// taskSrv := task.NewService(taskRepo)
	// taskEnd := task.MakeEnponints(taskSrv)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.POST("/login", gin.HandlerFunc(userEnd.LoginUser))
	router.POST("/register", gin.HandlerFunc(userEnd.RegisterUser))
	router.GET("/users", middleware.ValidToken, gin.HandlerFunc(userEnd.GetUser))

	router.Run(":8080")

}
