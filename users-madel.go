package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

func createUser(db *gorm.DB,user *User)  error{
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashPwd)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func loginUser(db *gorm.DB, user *User) (string, error) {
	//get user from email
	selectedUser := new(User)
	result := db.Where("email = ?", user.Email).First(&selectedUser)
	if result.Error != nil {
		fmt.Println("error getting user")
		return "", result.Error
	}

	//compare password
	err:= bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Println("error comparing password")
		return "", err
	}

	//return jwt token
	//create token
	jwtTest := "test"
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// sign token
	t, err := token.SignedString([]byte(jwtTest))
	if err != nil {
		fmt.Println("error signing token")
		return "", err
	}

	return t, nil

}