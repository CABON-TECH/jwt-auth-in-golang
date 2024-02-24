package controllers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"github.com/CABON-TECH/jwt-auth-in-golang/models"
	"github.com/CABON-TECH/jwt-auth-in-golang/initializers"

)

func Signup (c *gin.Context) {

	var body struct {
		Username string
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	db := initializers.DB

	user := models.User{
		Username: body.Username,
		Email: body.Email,
		Password: string(hash),
}
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}


	c.JSON(http.StatusOK, gin.H{})
}
