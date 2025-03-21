package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null;unique"`
	Avatar    string
	Status    bool   `gorm:"not null;default:true"`
	Roles     []Role `gorm:"many2many:user_roles"`
}
