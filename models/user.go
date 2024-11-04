// models/user.go
package models

type User struct {
    ID    int    `json:"id"`
    Nome  string `json:"nome"`
    Email string `json:"email"`
}
