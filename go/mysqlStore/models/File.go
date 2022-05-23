package model

import (
	"github.com/jinzhu/gorm"
)

type fileInfo struct {
	Name       string `gorm:"not null"`
	Content    []byte `gorm:"type:LONGBLOB;not null"`
	Size       int64  `gorm:"not null"`
	Filetype   string `gorm:"not null"`
	Source     string `gorm:"not null"`
	ClientIp   string `gorm:"not null"`
	RemoteBool bool
	CreateUser string
	UpdateUser string
	gorm.Model
}

func FileNew() *fileInfo {
	return &fileInfo{}
}
func (f *fileInfo) Insert() error {
	if err := db.Create(&f).Error; err != nil {
		return err
	}
	return nil
}

func (f *fileInfo) Delete() error {
	if err := db.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}
