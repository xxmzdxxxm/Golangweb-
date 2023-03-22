package dboption

import (
	"log"
	loverwaitdb "loverWait/loverwaitDB"
	"loverWait/model"
)

func CreateTableComment() {
	db := loverwaitdb.GetDB()
	err := db.AutoMigrate(model.Comments{})
	if err != nil {
		log.Fatal("create wrong")

	} else {
		log.Println("Create suceessed")
	}
}

func InsertComment(s string) {
	db := loverwaitdb.GetDB()
	com := model.Comments{
		Say: s,
	}
	db.Create(&com)
}
