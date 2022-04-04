package services

import (
	"blog/models"
	"time"
)

type User struct {
	TableName   string
	ID          string
	UID         string
	Name        string
	Password    string
	ArticleId   string
	RoleId      string
	Email       string
	Phone       string
	Address     string
	CreateTime  time.Time
	CreateUser  string
	UpdatedTime time.Time
	UpdatedUser string
}
type TableUser interface {
	GetTableName() string
}

func (u *User) GetTableName() string {
	return u.TableName
}

func (u *User) Add() error {
	user := map[string]interface{}{
		"TableName": u.TableName,
		"ID":        u.ID,
		"UID":       u.UID,
		"Name":      u.Name,
		"Password":  u.Password,
		// "articleId":   u.ArticleId,
		"RoleId":      u.RoleId,
		"Email":       u.Email,
		"Phone":       u.Phone,
		"Address":     u.Address,
		"CreateTime":  u.CreateTime,
		"UpdatedTime": u.UpdatedTime,
		"UpdatedUser": u.UpdatedUser,
		"CreateUser":  u.CreateUser,
	}
	if err := models.CreateUser(user); err != nil {
		return err
	}
	return nil
}
