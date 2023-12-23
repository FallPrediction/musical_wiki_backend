package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Actor struct {
	Id             uint32
	Name           string
	TranslatedName string
	NickName       string
	Nationality    string
	Born           string
	ImageId        *uint32
	Content        string
	Socials        datatypes.JSON
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
