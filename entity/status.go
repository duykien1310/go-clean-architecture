package entity

import "time"

type Status struct {
	Id   int `gorm:"primaryKey"`
	Code string
	Name string
	Type string

	CreatedAt time.Time
	UpdatedAt time.Time
}
