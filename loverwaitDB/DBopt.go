package loverwaitdb

import (
	"fmt"

	"gorm.io/gorm"
)

type DreamPlace struct {
	gorm.Model
	Provin string `gorm:"->;<-:create"`
	Name   string
}

func (DreamPlace) TableName() string { //绑定这个数据结构对应的表单名字
	return "dreamp"
}

func CreateTAB() {
	db := GetDB()
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&DreamPlace{})
	if err != nil {
		fmt.Println("create error")
	} else {
		println("cerate sucessed")
	}

}

func InsertDB(pro string, name string) {
	db := GetDB()
	dreamp := DreamPlace{
		Name:   name,
		Provin: pro,
	}
	db.Create(&dreamp)

}

func InsertDBGrop(tab []DreamPlace) {
	db := GetDB()
	db.Create(&tab)
}

func SelectByname(name string) *DreamPlace {
	db := GetDB()
	dre := DreamPlace{}
	db.Where("name = ?", name).Take(&dre)
	return &dre
}

func DeleteByname(name string) {
	db := GetDB()
	db.Delete(SelectByname(name))
}

func UpdateByname(name string, Nname string) {
	db := GetDB()
	dre := DreamPlace{}
	dre = *SelectByname(name)
	dre.Name = Nname
	db.Save(dre)
}
