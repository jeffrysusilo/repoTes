package models

type User struct {
    UserID    string `json:"user_id"`
    Name      string `json:"name"`
    Phone     string `json:"phone"`
    Email     string `json:"email"`
}
