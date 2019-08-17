package jwt

var jwtKey = []byte("my_secret_key")

token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "foo": "bar",
    "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
})

tokenString, err := token.SignedString(hmacSampleSecret)
