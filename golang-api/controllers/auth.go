package controllers

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/yourusername/go-ticket/config"
)

var JwtKey = []byte("secret_key")

func Login(w http.ResponseWriter, r *http.Request) {
    type LoginRequest struct {
        Email string `json:"email"`
    }
    var req LoginRequest
    json.NewDecoder(r.Body).Decode(&req)

    var userID string
    err := config.DB.QueryRow("SELECT user_id FROM users WHERE email=$1", req.Email).Scan(&userID)
    if err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, _ := token.SignedString(JwtKey)
    json.NewEncoder(w).Encode(map[string]string{
        "token": tokenString,
    })
}
