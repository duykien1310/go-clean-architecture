package entity

import (
	"time"
)

type Following struct {
	Id           int `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserId       int
	FollowUserId int

	CreatedAt time.Time
	UpdatedAt time.Time
}
