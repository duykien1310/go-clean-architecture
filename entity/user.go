package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserName  string `gorm:"type:varchar(255)" json:"userName"`
	Password  string `gorm:"type:varchar(255)" json:"password"`
	FirstName string `gorm:"type:varchar(255)" json:"firstName"`
	LastName  string `gorm:"type:varchar(255)" json:"lastName"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}
