package model

import "gorm.io/gorm"

type CompanyType struct {
    gorm.Model
    Identifier string
    Name string
}

func (CompanyType) TableName() string {
    return "m_company_types"
}
