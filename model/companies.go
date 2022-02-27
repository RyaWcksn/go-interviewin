package model

import "gorm.io/gorm"

type Company struct {
    gorm.Model
    Name string
    PersonInCharge string
    Address string
    CompanyType int
    FKCompanyType CompanyType `gorm:"foreignkey:CompanyType"`
    PhoneNumber string
}

func (Company) TableName() string {
    return "u_companies"
}
