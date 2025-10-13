package entity

import (
	"time"
)

type Post struct {
	Id      int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title   string `gorm:"type:text" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserId  int
	User    *User `gorm:"foreignKey:UserId"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
