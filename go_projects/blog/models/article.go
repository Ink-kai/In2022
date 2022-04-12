package models

import (
	"errors"
)

type Article struct {
	Uid     string `gorm:"type:varchar(255);UNIQUE_INDEX"`
	Title   string `gorm:"type:varchar(255);not null"`
	Content string `gorm:"not null"`
	BasicModel
}

type Articler interface {
	CreateArticle() error
	UpdateArticle(aid string) error
}

func ArticleNew(data Article) Articler {
	return &data
}

func (a *Article) CreateArticle() error {
	result := db.Create(&a)
	if result.RowsAffected != 0 {
		return errors.New("success")
	}
	return nil
}

func (a *Article) UpdateArticle(aid string) error {
	result := db.Create(&a)
	if result.RowsAffected != 0 {
		return errors.New("success")
	}
	return nil
}
