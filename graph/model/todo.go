package model

import (
	"gorm.io/gorm"
)

// type Todo struct {
// 	gorm.Model
// 	Text   string `json:"text"`
// 	Done   bool   `json:"done" gorm:"default:false"`
// 	UserID uint   `json:"userId" gorm:"index"`
// 	User   User
// }

type Todo struct {
	gorm.Model
	Text   string `json:"text"`
	Done   bool   `json:"done" gorm:"default:false"`
	UserID uint   `json:"userId" gorm:"index"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

// Gorm by default define foreignKey by User model & + ID so by default in this case
// is `UserID`. To overwrite this default define by string in struct. Why i put in there because i want Todo
// as learning perpuso
