package model

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Place string `gorm:"column:place"`
	Did   bool   `gorm:"column:did"`
	Who   string `gorm:"column:who"`
}

type Comments struct {
	ID  int
	Say string
}

func (Place) TableName() string {
	return "LoverPlace"
}

func (Comments) TableName() string {
	return "comment"
}
