package model

import "time"

// User model
type User struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	UserName   string    `json:"user_name"`
}
