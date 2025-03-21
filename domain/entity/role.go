package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Users       []User `gorm:"many2many:user_roles"`
}
