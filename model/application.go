package model

import "gorm.io/gorm"

type Application struct {
    gorm.Model
    InterviewerId int
    FKInterviewerId Interviewers `gorm:"ForeignKey:InterviewerId"`
    CandidateName string
    CompanyId int
    FKCompanyId Company `gorm:"ForeignKey:CompanyId"`
    CompanyName string
    JobTitle string
    Status string
}

func (Application) TableName() string {
    return "u_applications"
}
