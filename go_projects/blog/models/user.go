package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uint64 `gorm:"primaryKey;AUTO_INCREMENT"`
	Uid      string `gorm:"primaryKey;index:idx_member"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string
	Phone    string
	State    int `gorm:"default:1;not null"`
	BasicModel
}

func uuid_generate_v3() string {
	slat := []byte(time.Now().UTC().String())
	h := md5.New()
	h.Write(slat)
	return hex.EncodeToString(h.Sum(nil))
}
func CreateUser(data map[string]interface{}) error {
	// fmt.Fprintf(os.Stdout, "%v\n", data)
	u := User{
		Uid:      uuid_generate_v3(),
		Name:     data["Name"].(string),
		Password: data["Password"].(string),
		Email:    data["Email"].(string),
		Phone:    data["Phone"].(string),
		// Model: BasicModel{
		// 	Created:     data["CreateTime"].(time.Time),
		// 	Updated:     data["UpdatedTime"].(time.Time),
		// 	UpdatedUser: data["UpdatedUser"].(string),
		// 	CreatedUser: data["CreateUser"].(string),
		// },
	}
	result := db.Create(&u)
	fmt.Fprintf(os.Stdout, "%v", result.RowsAffected)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateUser(uid string, data interface{}) error {
	if e := db.Model(&User{}).Where("uid=? and state=?", uid, true).Updates(data).Error; e != nil {
		return e
	}
	return nil
}

func DeleteUser(uid string) error {
	if e := db.Where("uid=? and state=?", uid, true).Delete(&User{}).Error; e != nil {
		return e
	}
	return nil
}

func GetUser(uid string) (*User, error) {
	var u User
	err := db.Where("uid=? and state=?", uid, true).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	err = db.Model(&u).Association("article").Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &u, err
}
