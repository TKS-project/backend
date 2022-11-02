package domain

import "time"

type Prefecture struct {
	ID        int `gorm:"primary_key auto_increment"`
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Citie struct {
	ID           int16 `gorm:"primary_key auto_increment"`
	PrefectureId int
	Name         string
	Code         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
