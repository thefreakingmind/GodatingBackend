package main

import (
	"fmt"
	"log"
	"os"
	"golang.org/x/crypto/ccrypt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

type User struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
}

// Connecting to the DB
func ConnectDB() *gorm.DB{
  err := godotenv.Load()
  if err != nil{
	log.Fatal("Error in Loading Page")
  }
  username := os.Getenv("db_user")
  password := os.Getenv("db_password")
  dbname := os.Getenv("db_name")
  dbHost := os.Getenv("db_host")
  dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbname, password) //Build connection string
  fmt.Println(dbUri)
  db, err :=gorm.Open("postgres", dbUri)

  if err != nil{
	fmt.Println("Error")
  }
  defer db.Close()
  db.AutoMigrate(User{})
  fmt.Println("SuccessFully Connected")
  return db
}

//Creating The User
func CreateUser(w http.ResponseWriter, r *http.Request){
  user := &User
  json.NewDecoder(r.body).Decode(user)
  pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err!=nil{
	fmt.Println(err)
	err:= ErrorResponse{
	  Err: "Password Encryption Failed"
	}
	json.NewEncoder(w).Encode(err)
  }
  user.Password = string(pass)
  crateduser := db.Create(user)
  var errMessage = createdUser.Error
  if createdUser.Error != nil{
	fmt.Println("Error")
  }
  json.NewEncoder(w).Encode(createduser)
}

//Login 
func Login(w http.ResponseWriter, r *http.Request) {
  user := &User{}
  err := json.NewDecoder(r.Body).Decode(user)
  if err != nil {
	var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
	json.NewEncoder(w).Encode(resp)
	return
  }
  resp := FindOne(user.Email, user.Password)
  json.NewEncoder(w).Encode(resp)
}


