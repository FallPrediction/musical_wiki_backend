package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Actor struct {
	Id             uint32
	Name           string
	TranslatedName string
	NickName       string
	Nationality    string
	Born           string
	Content        string
	Credits        []Credit
	Avatar         string
	Socials        datatypes.JSON
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
