package main

import "api-jwt-go/models"

func init() {
	models.LoadEnvs()
	models.ConnectDB()
}

func main() {
	models.DB.AutoMigrate(&models.User{})
}
