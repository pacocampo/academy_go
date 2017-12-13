package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func Auth(w http.ResponseWriter, req *http.Request) {
	login := Login{}

	json.NewDecoder(req.Body).Decode(&login)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": login.Username,
	})
	tokenString, error := token.SignedString([]byte("thisismysecretword"))
	if error != nil {
		fmt.Println(error)
	}
	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func Validate(value string) (bool, string) {
	token, _ := jwt.Parse(value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("thisismysecretword"), nil
	})

	if token == nil {
		return false, "Invalid authorization token"
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var login Login
		mapstructure.Decode(claims, &login)
		return true, login.Username
	}
	return false, "Invalid authorization token"
}
