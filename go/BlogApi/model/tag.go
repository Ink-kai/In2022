package model

type Tag struct {
	ID        uint `gorm:"primary_key;autoIncrement:true"`
	Tid       uint `gorm:"primaryKey;autoIncrement:false"`
	ArticleID uint
	Title     string `gorm:"not null"	form:"not null"`
	Desc      string `gorm:"not null"	form:"not null"`
	CreatedAt uint
	UpdatedAt uint
	DeletedAt uint `sql:"index"`
}
