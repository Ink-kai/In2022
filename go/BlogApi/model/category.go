package model

type Category struct {
	ID        uint `gorm:"primary_key;autoIncrement:true"`
	Cid       uint `gorm:"primaryKey;autoIncrement:false"`
	ArticleID uint
	Title     string `gorm:"not null"	form:"not null"`
	Desc      string `gorm:"not null"	form:"not null"`
	CreatedAt uint
	UpdatedAt uint
	DeletedAt uint `sql:"index"`
}
