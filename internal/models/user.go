package models

type User struct {
    Name     string   `json:"name"`
    Email    string   `json:"email"`
    Groups   []string `json:"groups"`
    Verified bool     `json:"verified"`
}
