package utility

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var envFilePath = ".env"

var jwtKey = DotEnv("TOKEN_KEY", envFilePath)
var authDuration, _ = strconv.Atoi(DotEnv("AUTH_DURATION", ".env"))

func GenerateJWTToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * time.Duration(authDuration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return signedToken, nil
}
