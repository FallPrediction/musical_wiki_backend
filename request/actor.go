package request

import (
	"gorm.io/datatypes"
)

type Actor struct {
	Name           string         `json:"name" binding:"required"`
	TranslatedName string         `json:"translated_name"`
	NickName       string         `json:"nick_name"`
	Nationality    string         `json:"nationality"`
	Born           string         `json:"born"`
	ImageId        *uint32        `json:"image_id"`
	Content        string         `json:"content"`
	Socials        datatypes.JSON `json:"socials" binding:"omitempty,json"`
}
