package models

import "time"

type Credit struct {
	Id        uint32
	Time      string
	Place     string
	Character string
	Musical   string
	ActorId   uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
