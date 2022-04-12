package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	Uid     string `gorm:"type:varchar(255);unique"`
	Account int    `gorm:"type:int;not null;unique"`
	// Articles []Article `gorm:"foreignKey:Uid"`
	Name     string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255)"`
	Phone    string `gorm:"type:varchar(255)"`
	Address  string `gorm:"type:varchar(255)"`
	State    int    `gorm:"type:int;default:1"`
	BasicModel
}

type User interface {
	CreateUser() error
	UpdateUser(uid string) error
}

func UserNew(data UserInfo) User {
	return &data
}

func (u *UserInfo) CreateUser() error {
	result := db.Create(&u)
	if result.RowsAffected != 0 {
		return errors.New("success")
	}
	return nil
}

func (u *UserInfo) UpdateUser(uid string) error {
	result := db.Model(&UserInfo{}).Where("uid=? and state=?", uid, 1).Updates(u)
	if result.RowsAffected != 0 && result.Error != gorm.ErrRecordNotFound {
		return errors.New("success")
	}
	return nil
}

func DeleteUser(uid string) error {
	result := db.Where("uid=? and state=?", uid, 1).Delete(&UserInfo{})
	if result.RowsAffected != 0 {
		return errors.New("success")
	}
	return nil
}

func GetUser(uid string) (*UserInfo, error) {
	var user UserInfo
	result := db.Where("uid=? and state=?", uid, 1).First(&user)
	if result.RowsAffected != 0 && result.Error != gorm.ErrRecordNotFound {
		return &user, errors.New("success")
	}
	return nil, errors.New("fail")
}

func GetAllUser() (*[]UserInfo, error) {
	var user []UserInfo
	result := db.Where("state=?", 1).Find(&user)
	if result.RowsAffected != 0 && result.Error != gorm.ErrRecordNotFound {
		return &user, errors.New("success")
	}
	return nil, errors.New("fail")
}
