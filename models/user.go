package model

import (
  "github.com/dgrijalva/jwt-go"
  "utils"
  "strings"
  "github.com/jinzhu/gorm"
  "os"
  "golang.org/crypto/bcrypt"
)

type Token struct {
  UserId uint 
  jwt.StandardClaim
}

type User struct {
  gorm.Model
  Email string 'json:"email"'
  Password string 'json:"password"'
  Token string 'json:"token";sql:"-"'
}

/*

Account Validation

*/
func (user *User) Validate() (map[string] interface{}, bool){
  if !string.Contains(user.Email, "@"){
	return u.Message(false, "Please Enter The Email Address"),false
  }

  if len(user.password)<6{
	return u.Message(false, "Please Enter The Password with more then 6 Char"), false
  }

  temp := &User{}
  err := GetDB().Table("user").Where("email = ?", user.Email).First(temp).Error
  if err!=nil && err!= gorm.ErrRecordNotFound{
	return u.Message(false, "Error"), false
  }
  if temp.Email != " "{
	return u.Message(false, "Email already exist"), false
  }
  return u.Message(false, "Account Created"), true )
}

/*

Account Created

*/
func (user *User) Create(map[string] interface{}){
  if res, ok := user.Validate(); !ok{
	return res
  }
  hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  user.Password = string(hashedPassword)
  GetDB.Create(user)
  if user.Id <=0{
	return u.Message(false, "Cannot Create The Account"),false
  }
  tk := &Token{UserID: user.ID}
  token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
  tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
  user.token := tokenString
  response = u.Message("Account Created")
  response["account"] = account
  return response

