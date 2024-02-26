package initializers

import ( 
	"github.com/CABON-TECH/jwt-auth-in-golang/models"
	"gorm.io/gorm"
	
)

func SyncDatabase (DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
}
