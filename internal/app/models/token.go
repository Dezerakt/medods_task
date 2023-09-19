package models

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	UserId uint `gorm:"column:user_id"`
}

func (t *Token) model() {

}
