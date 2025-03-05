package services

import (
	"online-Market_project_Golang-Backent/internal/db"
	"online-Market_project_Golang-Backent/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user models.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)
	if err := db.DB.Create(&user).Error; err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
	})
	tokenString, err := token.SignedString([]byte("qwertyuiop"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
