package model

import "gorm.io/gorm"

type FileType struct {
    gorm.Model
    Name string
    Description string
}

func (FileType) TableName() string {
    return "m_file_types"
}
