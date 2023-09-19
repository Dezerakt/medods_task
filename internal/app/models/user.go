package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	GUID string `gorm:"column:guid; size:255;"`

	Token Token `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *User) model() {

}
