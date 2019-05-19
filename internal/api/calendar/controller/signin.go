package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/roger-russel/deploy-calendar/pkg/api"
	"github.com/valyala/fasthttp"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(os.Getenv("JWT-KEY"))

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

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}

	api.JsonOut(ctx, struct {
		Token string `json:"token"`
	}{Token: tokenString})

}
