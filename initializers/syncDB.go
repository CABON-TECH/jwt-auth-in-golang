package initializers

import ( 
	"github.com/CABON-TECH/jwt-auth-in-golang/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SyncDatabase () {
	DB.AutoMigrate(&models.User{})
}
