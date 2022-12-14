package domain

import (
	"time"
)

type Sightseeing struct {
	PrefectureId int    `json:"prefecture_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Description  string `json:"description"`
	City         string `json:"city"`
	UpdatedAt    time.Time
	CreatedAt    time.Time
	DeletedAt    time.Time
}
