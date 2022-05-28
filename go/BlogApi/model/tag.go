package model

type Tag struct {
	ID        uint64          `gorm:"primary_key;autoIncrement:true"`
	TagID     uint64          `gorm:"primaryKey;autoIncrement:false"`
	Tags      []ArticleRefTag `gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title     string          `gorm:"not null"	form:"not null"`
	Desc      string          `gorm:"not null"	form:"not null"`
	CreatedAt uint64
	UpdatedAt uint64
	DeletedAt uint64 `sql:"index"`
}
