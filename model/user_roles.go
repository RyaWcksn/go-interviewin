package model

import "gorm.io/gorm"

type UserRole struct {
    gorm.Model
    Identifier string
    Name string
}

func (UserRole) TableName() string {
    return "m_user_roles"
}
