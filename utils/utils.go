package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	Password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(Password), err
}

func PasswordMatch(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func GenerateJWT(username string) string {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      jwt.TimeFunc().Add(time.Hour * 24).Unix(),
	})
	token, _ := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}
func ParseToken(token string) (string, error) {
	tokenVal, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := tokenVal.Claims.(jwt.MapClaims); ok && tokenVal.Valid {
		return claims["username"].(string), nil
	}
	return "", nil
}
