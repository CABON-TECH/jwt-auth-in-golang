package main

import (
	"github.com/gin-gonic/gin"

	"github.com/CABON-TECH/jwt-auth-in-golang/controllers"
	"github.com/CABON-TECH/jwt-auth-in-golang/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabase()

}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", controllers.Validate)
	r.Run()
}
