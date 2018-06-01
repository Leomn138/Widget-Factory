package handler

import (
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
	"time"
	"github.com/Leomn138/Widget-Factory/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

func CreateToken(auth *config.Auth, w http.ResponseWriter, req *http.Request) {
	var credentials Credentials
	error := json.NewDecoder(req.Body).Decode(&credentials)
	if error != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": credentials.Username,
		"password": credentials.Password,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, error := token.SignedString([]byte(auth.Secret))
	if error != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
}

func ValidateMiddleware(auth *config.Auth, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(auth.Secret), nil
				})
				if error != nil {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
			}
		} else {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	})
}