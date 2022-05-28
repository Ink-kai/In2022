package model

type Article struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true"`
	Aid       uint `gorm:"primaryKey;autoIncrement:false"`
	UserID    uint
	Categorys []Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Tags      []Tag      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title     string     `gorm:"not null"	form:"not null"`
	Content   string     `gorm:"not null"	form:"not null"`

	CreatedAt uint
	UpdatedAt uint
	DeletedAt uint `sql:"index"`
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
