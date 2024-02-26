package initializers

import (
	"github.com/CABON-TECH/jwt-auth-in-golang/models"
)

//var DB *gorm.DB

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
