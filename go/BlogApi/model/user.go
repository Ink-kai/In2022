package model

type User struct {
	ID        uint `gorm:"primary_key;autoIncrement:true"`
	Uid       uint `gorm:"primaryKey;autoIncrement:false"`
	ArticleID uint
	Articles  []Article `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name      string    `gorm:"not null"	form:"not null"`
	Password  string    `gorm:"not null"`
	Email     string
	Phone     string
	Birthday  uint
	Address   string
	Desc      string `gorm:"not null"	form:"not null"`
	Remark    string
	CreatedAt uint
	UpdatedAt uint
	DeletedAt uint `sql:"index"`
}

func (user User) GetUser(uid string) error {
	if uid == "" {
		if err := DB.Joins("user").Find(&user).Error; err != nil {
			return err
		}
	} else {
		if err := DB.Joins("user").Where("uid like ?", "%"+uid+"%").Find(&user).Error; err != nil {
			return err
		}
	}
	return nil
}

func (user User) AddUser() error {
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (user User) DelUser() error {
	if err := DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (user User) UpdateUser() error {
	if err := DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
