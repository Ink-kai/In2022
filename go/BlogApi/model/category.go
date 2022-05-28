package model

type Category struct {
	ID         uint64               `gorm:"primary_key;autoIncrement:true"`
	CategoryID uint64               `gorm:"primaryKey;autoIncrement:false"`
	Categorys  []ArticleRefCategory `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title      string               `gorm:"not null"	form:"not null"`
	Desc       string               `gorm:"not null"	form:"not null"`
	CreatedAt  uint64
	UpdatedAt  uint64
	DeletedAt  uint64 `sql:"index"`
}
