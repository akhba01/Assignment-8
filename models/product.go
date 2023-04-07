package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `json:"title" form:"title" valid:"required~Title of product is required"`
	Description string `json:"description" form:"description" valid:"required~Description of Product is required"`
	UserID      uint
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCre := govalidator.ValidateStruct(p)

	if errCre != nil {
		err = errCre
		return
	}

	err = nil
	return

}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCre := govalidator.ValidateStruct(p)

	if errCre != nil {
		err = errCre
		return
	}

	err = nil
	return

}
