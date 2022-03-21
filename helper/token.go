package helper

import (
	"errors"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET_KEY = []byte(os.Getenv("SECRECT_KEY"))

type LoginDetails struct {
	Uid string
	jwt.StandardClaims
}

func GenerateTokens(user_id string) (Token string, RefreshToken string, err error) {
	claims := &LoginDetails{
		Uid: user_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(12)).Unix(),
		},
	}

	refreshClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(48)).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(SECRET_KEY)
	if err != nil {
		log.Panic(err)
		return
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(SECRET_KEY)
	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func CheckToken(token string) (Logins *LoginDetails, err error) {
	t, err := jwt.ParseWithClaims(
		token,
		&LoginDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return SECRET_KEY, nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := t.Claims.(*LoginDetails)
	if !ok {
		err = errors.New("the token is invalid")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("the token has expired")
		return
	}

	return claims, err
}
