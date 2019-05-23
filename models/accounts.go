package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	u "goapi/utils"
	"strings"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account

type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token" sql:"-"`
}

// Validate incoming user details...

func (account *Account) Validate() (map[string] interface{}, bool){
	fmt.Println("validation beginning!");
	if !strings.Contains(account.Email, "@"){
		return u.Message(false, "Email address is required"), false
	}
	fmt.Println("email address is good")
	if len(account.Password) < 6{
		return u.Message(false, "Password is required and must be 6 word"), false
	}
	//email must be unique
	temp := &Account{}

	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	fmt.Println("record found!")
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}
	fmt.Println("record found!")
	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string] interface{}){
	if resp, ok := account.Validate(); !ok {
		return resp
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account) // neglect this code will be disaster hell!!

	if(account.ID <= 0) { //check whether the account is created!
		return u.Message(false, "Failed to create account, connection error.")
	}
	//create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password after used

	response := u.Message(true, "Account has been create")
	response["account"] = account
	return response
}

func Login(email, password string)(map[string] interface{}){
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not Found")
		}
		return u.Message(false, "Connection error. please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		//password not match!
		return u.Message(false, "Invalid login credentials. please try again")
	}
	//worked! Logged In

	//create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString // store the token in the response

	resp := u.Message(true, "Logged in")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {
	acc := &Account{}
	GetDB().Table("accounts").Where("id=?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}
	acc.Password = ""
	return acc
}