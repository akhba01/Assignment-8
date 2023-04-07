package models

import (
	"assignment8/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" form:"full_name" json:"full_name" valid:"required~Full Name is required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Your email is invalid format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlenth(6)~Password must be have a minimun length of 6 character"`
	Role     string    `gorm:"not null" form:"role" json:"role" validate:"required"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCre := govalidator.ValidateStruct(u)

	if errCre != nil {
		err = errCre
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return nil
}
