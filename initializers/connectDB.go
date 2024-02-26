package initializers

import (
	"os"
	"gorm.io/driver/postgres"
        "gorm.io/gorm"
	
  )
/*
  const (
	  host     = "localhost"
	  port     = 5432
	  user     = "postgres"
	  password = "cabon"
	  dbname   = "jwt_db"
  )
  var DB *gorm.DB


  func ConnectDB () {
	  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	   host, port, user, password, dbname)
	   DB, err := gorm.Open("postgres", psqlInfo)
	  if err != nil {
		  panic(err)
	  }

	  err = DB.Ping()
	  if err != nil {
		  panic(err)
	  }
	  fmt.Println("Successfully connected!")
	  return DB
  }*/

var DB *gorm.DB
func connectDB () {
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}
}
