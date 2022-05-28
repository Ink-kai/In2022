package model

type User struct {
	ID        uint64           `gorm:"primary_key;autoIncrement:true"`
	UserUID   uint64           `gorm:"primaryKey;autoIncrement:false;unique"`
	Users     []UserRefArticle `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name      string           `gorm:"not null"	form:"Name"`
	Password  string           `gorm:"not null"	form:"Password"`
	Email     string           `form:"Email"`
	Phone     string           `form:"Phone"`
	Birthday  int64            `form:"Birthday" validate:"email"`
	Address   string           `form:"Address"`
	Desc      string           `gorm:"type:varchar(4000)"	form:"Desc"`
	Remark    string           `gorm:"type:varchar(4000)"	form:"Remark"`
	CreatedAt uint64
	UpdatedAt uint64
	DeletedAt uint64 `sql:"index"`
}
type UserRefArticle struct {
	ID        uint64 `gorm:"primary_key;autoIncrement:true"`
	ArticleID uint64
	UserID    uint64
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
