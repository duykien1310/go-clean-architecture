package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"
)

type User struct {
	Id          int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserName    string `gorm:"type:varchar(255)" json:"userName"`
	Password    string `gorm:"type:varchar(255)" json:"password"`
	FirstName   string `gorm:"type:varchar(255)" json:"firstName"`
	LastName    string `gorm:"type:varchar(255)" json:"lastName"`
	Email       string `gorm:"type:varchar(255)" json:"email"`
	PhoneNumber string `gorm:"type:varchar(255)" json:"phoneNumber"`
	Address     string `gorm:"type:varchar(255)" json:"address"`
	RoleId      int
	Role        *Role `gorm:"foreignKey:RoleId"`
	StatusId    int
	Status      *Status
	IsDelete    bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (n *User) normalizeInformation() {
	n.FirstName = norm.NFC.String(n.FirstName)
	n.LastName = norm.NFC.String(n.LastName)
	n.Email = norm.NFC.String(n.Email)
	n.UserName = norm.NFC.String(n.UserName)
}

func (n *User) BeforeUpdate(tx *gorm.DB) (err error) {
	n.normalizeInformation()

	return nil
}

func (n *User) BeforeCreate(tx *gorm.DB) (err error) {
	n.normalizeInformation()

	return nil
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
