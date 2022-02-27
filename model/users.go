package model

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name string `gorm:"type:varchar(100);not null"`
    Email string `gorm:"type:varchar(100);not null"`
    Password string `gorm:"type:varchar(100);not null"`
    Role int `gorm:"type:int;not null"`
}

func (User) TableName() string {
    return "users"
}
