package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username  string `gorm:"type:varchar(100);not null"`
    Email     string `gorm:"type:varchar(100);unique;not null"`
    Password  string `gorm:"type:varchar(100);not null"`
    IsActive  bool   `gorm:"default:true"`
    CompanyID uint   // Foreign key for Company
    Company   Company
}

