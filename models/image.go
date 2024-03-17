package models

import (
	"time"
)

type Image struct {
	Id        uint32
	ImageName string
	Mime      string
	ActorId   uint32
	ImageType string
	CreatedAt time.Time
	UpdatedAt time.Time
}
