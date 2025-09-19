package middleware

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("secret_key") 

func JwtVerify(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing auth token", http.StatusUnauthorized)
            return
        }

        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return JwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
