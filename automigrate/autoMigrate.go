package automigrate

import "api-jwt-go/models"

func AutoMigrate() {
	models.DB.AutoMigrate(&models.User{})
}
