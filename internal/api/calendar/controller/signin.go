package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/valyala/fasthttp"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Claims Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//Signin function
func Signin(ctx *fasthttp.RequestCtx) {

	var creds = Credentials{}

	err := json.Unmarshal(ctx.PostBody(), &creds)

	if err != nil {
		ctx.SetStatusCode(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		ctx.SetStatusCode(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json; charset=utf-8")

	fmt.Fprintf(ctx, tokenString, &ctx.Request)

}
