package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Minifyr struct {
	ID			uint64	`json:"id" gorm:"primaryKey"`
	Redirect	string	`json:"redirect" gorm:"not null"`
	Minifyr		string	`json:"minifyr" gorm:"unique;not null"`
	Clicked		uint64	`json:"clicked"`
	Random		bool	`json:"random"`
}


func Setup() {

	dataSource := "host=0.0.0.0 user=postgres password=minifyrtest123 dbname=Minifyr port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Minifyr{})
	if err != nil {
		fmt.Println(err)
	}
}
