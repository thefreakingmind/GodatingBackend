package main

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "time"
  "log"
  "os"
  _"strings"
  _"context"
  "net/http"
  "golang.org/x/crypto/bcrypt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/joho/godotenv"
  "github.com/dgrijalva/jwt-go"
)
type Exception struct{
  Message string `json:"message"`
}

type User struct {
  gorm.Model
  Email string `json:"email"`
  Password string `json:"password"`
}


type Claims struct {
  UserID uint
  Email  string
  *jwt.StandardClaims
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

var db = ConnectDB()

type ErrorResponse struct {
  Err string
}

type error interface {
  Error() string
}

//Creating The User
func CreateUser(w http.ResponseWriter, r *http.Request){
  user := &User{}
  json.NewDecoder(r.Body).Decode(user)
  pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err!=nil{
  fmt.Println(err)
  err:= ErrorResponse{
    Err: "Password Encryption Failed",
  }
  json.NewEncoder(w).Encode(err)
  }
  user.Password = string(pass)
  createdUser := db.Create(user)
  //var errMessage = createdUser.Error
  if createdUser.Error != nil{
  fmt.Println("Error")
  }
  json.NewEncoder(w).Encode(createdUser)
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
  resp := Find(user.Email, user.Password)
  json.NewEncoder(w).Encode(resp)
}


func Find(email, password string) map[string]interface{} {
  user := &User{}
  if err := db.Where("Email = ?", email).First(user).Error; err != nil {
  var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
  return resp
  }
  expiresAt := time.Now().Add(time.Minute * 100000).Unix()
  errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
  if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
  var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
  return resp
  }
  tk := &Claims{
    UserID: user.ID,
    Email:  user.Email,
    StandardClaims: &jwt.StandardClaims{
      ExpiresAt: expiresAt,
      },
    }
  token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
  tokenString, error := token.SignedString([]byte("secret"))
  if error != nil {
  fmt.Println(error)
  }
  var resp = map[string]interface{}{"status": false, "message": "logged in"}
  resp["token"] = tokenString //Store the token in the response
  resp["user"] = user
  return resp
}


func main(){
  router := mux.NewRouter()
  router.HandleFunc("/register", CreateUser).Methods("POST")
  router.HandleFunc("/login", Login).Methods("POST")
  log.Fatal(http.ListenAndServe(":8000", router))
}
