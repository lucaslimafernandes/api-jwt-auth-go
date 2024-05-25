package main

import (
	"api-jwt-go/controllers"
	"api-jwt-go/models"

	"github.com/gin-gonic/gin"
)

func init() {
	models.LoadEnvs()
	models.ConnectDB()
}

func main() {

	router := gin.Default()

	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	// router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	router.Run()

}
