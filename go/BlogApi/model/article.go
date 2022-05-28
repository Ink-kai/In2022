package model

type Article struct {
	ID        uint64               `gorm:"primaryKey;autoIncrement:true"`
	ArticleID uint64               `gorm:"primaryKey;autoIncrement:false"`
	Articles  []UserRefArticle     `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Categorys []ArticleRefCategory `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tags      []ArticleRefTag      `gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title     string               `gorm:"not null"	form:"not null"`
	Content   string               `gorm:"not null"	form:"not null"`
	CreatedAt uint64
	UpdatedAt uint64
	DeletedAt uint64 `sql:"index"`
}
type ArticleRefCategory struct {
	ID         uint64 `gorm:"primary_key;autoIncrement:true"`
	ArticleID  uint64
	CategoryID uint64
}

type ArticleRefTag struct {
	ID        uint64 `gorm:"primary_key;autoIncrement:true"`
	ArticleID uint64
	TagID     uint64
}

func (article *Article) GetArticle(title string) error {
	if title == "" {
		if err := DB.Joins("user").Find(&article).Error; err != nil {
			return err
		}
	} else {
		if err := DB.Joins("user").Where("title like ?", "%"+title+"%").Find(&article).Error; err != nil {
			return err
		}
	}
	return nil
}

func (article *Article) AddArticle() error {
	if err := DB.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func (article *Article) DelArticle() error {
	if err := DB.Delete(&article).Error; err != nil {
		return err
	}
	return nil
}

func (article *Article) UpdateArticle() error {
	if err := DB.Save(&article).Error; err != nil {
		return err

	}
	return nil
}
