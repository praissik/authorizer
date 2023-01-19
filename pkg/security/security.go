package security

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func GenerateBcryptHash(text string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), 0)
	if err != nil {
		log.Panic(err)
	}
	return string(hash)
}

func CompareHashAndText(hash string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if err != nil {
		return false
	}
	return true
}

func NewToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"time":  time.Now().UTC().Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return "", fmt.Errorf("try again")
	}

	return tokenString, nil
}

func ValidToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return os.Getenv("APP_SECRET"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["UserEmail"], claims["Time"])
		return nil
	}
	return err
}
