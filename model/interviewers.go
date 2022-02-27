package model

import "gorm.io/gorm"

type Interviewers struct {
    gorm.Model
    UserId int
    FKUserId User `gorm:"foreignkey:UserId"`
    Name string
    PhoneNumber string
    CvFile int
    FKCvFile FilesUpload `gorm:"foreignkey:CvFile"`
    VideoFile int
    FKVideoFile FilesUpload `gorm:"foreignkey:VideoFile"`
}

func (Interviewers) TableName() string {
    return "u_interviewers"
}
