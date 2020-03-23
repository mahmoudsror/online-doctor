package models

import (
	"time"
)

type User struct {
	ID         uint32 `json:"id"`
	Fname      string `json:"fname"`
	Lname      string `json:"lname"`
	Email      string `json:"email"`
	Mobile     string `json:"mobile"`
	Password   string `json:"password"`
	UserType   string `json:"userType"`
	Creator    uint32
	Active     bool      `json:"active"`
	Verified   bool      `json:"verified"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
