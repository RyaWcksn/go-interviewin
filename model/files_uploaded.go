package model

import "gorm.io/gorm"

type FilesUpload struct {
    gorm.Model
    UserId int
    FKUserId User `gorm:"foreignkey:UserId"`
    FileType int
    FKFileType FileType `gorm:"foreignkey:FileType"`
    Name string
    Path string
}

func (FilesUpload) TableName() string {
    return "files_upload"
}
