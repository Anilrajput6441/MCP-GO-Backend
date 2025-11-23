package models

import "time"

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"password"`
	FullName  string    `json:"full_name" bson:"full_name"`
	Role      string    `json:"role" bson:"role"` // "user" or "admin"
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
