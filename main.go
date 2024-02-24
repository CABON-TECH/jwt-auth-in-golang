package main

import (
	
	"github.com/gin-gonic/gin"

	"github.com/CABON-TECH/jwt-auth-in-golang/initializers"
	 "github.com/CABON-TECH/jwt-auth-in-golang/controllers"
)
func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabase()

	
}

func main () {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.Run() 
}
