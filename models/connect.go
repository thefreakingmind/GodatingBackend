package models

import (
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/jinzhu/gorm"
  "os"
  "fmt"
  "github.com/joho/godotenv"
)

var db *gorm.DB

func init(){
  e := godotenv.Load()
  if e!= nil{
	fmt.Print("Error")
  }
  username := os.Getenv("db_user")
  password := os.Getenv("db_password")
  dbname := os.Getenv("db_name")
  dbHost := os.Getenv("db_host")
  dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbname, password) //Build connection string
  fmt.Println(dbUri)
  conn, err :=gorm.Open("postgres", dbUri)

  if err!=nil{
	fmt.Println("Error")
  }
  db = conn
  db.Debug().AutoMigrate(&User{})
}

func startDB() *gorm.DB{
  return db
}
