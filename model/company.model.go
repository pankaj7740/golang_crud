package model

import "gorm.io/gorm"

type Company struct {
    gorm.Model
    Name  string  `gorm:"type:varchar(100);not null"`
    GSTNo string  `gorm:"type:varchar(50);not null"`
    Users []User  `gorm:"foreignKey:CompanyID"`
}