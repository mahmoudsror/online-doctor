package models

import (
	"fmt"
	"time"

	"github.com/mahmoudsror/online-doctor/connections"
)

type User struct {
	ID         uint32    `json:"id"`
	Fname      string    `json:"fname"`
	Lname      string    `json:"lname"`
	Email      string    `json:"email"`
	Mobile     string    `json:"mobile"`
	Password   string    `json:"password"`
	Usertype   string    `json:"usertype"`
	Creator    uint32    `json:"creator"`
	Active     bool      `json:"active"`
	Verified   bool      `json:"verified"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (user *User) CreateUser() {
	fmt.Println("user data ", user)
	connections.GetDB().Create(user)
}
