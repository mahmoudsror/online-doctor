package models

import "time"

type User struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Fname     string `gorm:"size:50;not null" json:"fname"`
	Lname     string `gorm:"size:50;not null" json:"lname"`
	Email     string `gorm:"type:varchar(100);not null;unique" json:"email"`
	Mobile    string `gorm:"size:15;not null;unique" json:"mobile"`
	Password  string `gorm:"size:255;not null" json:"password"`
	UserType  string `gorm:"size:20;not null" json:"userType"`
	creator   uint32
	active    bool      `gorm:"default:false"`
	verified  bool      `gorm:"default:false"`
	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
