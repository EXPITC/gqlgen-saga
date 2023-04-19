package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Todo []Todo `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:ID"`
}

// `gorm:"foreignKey:UserID,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
// When you assign Todo to a user,
// GORM will save the user’s ID into Todo’ UserID field.

// Actually form take default model name + ID as default but this for example purpose
// you can define it manually for the overide purpose or just make sure, well its a options
// i choose to use it as reference
