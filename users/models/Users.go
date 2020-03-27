package models

import (
	"fmt"
	"time"

	"github.com/mahmoudsror/online-doctor/connections"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Fname      string    `gorm:"size:30" json:"fname"`
	Lname      string    `gorm:"size:30;" json:"lname"`
	Email      string    `gorm:"size:100;not null;unique" json:"email"`
	Mobile     string    `gorm:"size:20;" json:"mobile"`
	Password   string    `gorm:"size:100;not null;" json:"password"`
	Usertype   string    `gorm:"size:10;not null;" json:"usertype"`
	Creator    uint32    `json:"creator"`
	Active     bool      `gorm:"default:true" json:"active"`
	Verified   bool      `gorm:"default:true" json:"verified"`
	Created_at time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Updated_at time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (user *User) CreateUser() {
	fmt.Println("user data ", user)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	connections.GetDB().Create(user)
	if user.ID <= 0 {
		fmt.Println("error : ", user)
	}
	//return createdUser
}
