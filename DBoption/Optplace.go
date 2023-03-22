package dboption

import (
	"log"
	loverwaitdb "loverWait/loverwaitDB"
	"loverWait/model"
)

func CreateTable() {
	db := loverwaitdb.GetDB()
	err := db.AutoMigrate(model.Place{})
	if err != nil {
		log.Fatal("create wrong")

	} else {
		log.Println("Create suceessed")
	}
}

func Insertplace(place string, is bool) {
	db := loverwaitdb.GetDB()
	pla := model.Place{
		Place: place,
		Did:   is,
	}
	db.Create(&pla)
}

func SelectByplace(name string) *model.Place {
	db := loverwaitdb.GetDB()
	Pla := model.Place{}
	db.Limit(1).Where("ID = ?", name).Take(&Pla)
	return &Pla
}

func SelectByID(name int) *model.Place {
	db := loverwaitdb.GetDB()
	Pla := model.Place{}
	db.Limit(1).Where("ID = ?", name).Take(&Pla)
	return &Pla
}

func Selectbybool(did bool) *[]model.Place {
	db := loverwaitdb.GetDB()
	var list []model.Place
	db.Where("Did = ?", did).Find(&list)
	return &list
}
