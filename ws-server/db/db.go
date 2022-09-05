package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DbConnection *gorm.DB

type UserSession struct {
	gorm.Model
	SessionMap []UserSessionEntry
}

type UserSessionEntry struct {
	gorm.Model
	UserSessionId uint
	Key           string
	Value         string
}

type Dummy struct {
	gorm.Model
	Value string
}

func Init() {
	conn, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln("Could not create or open database!")
	}

	DbConnection = conn

	err = conn.AutoMigrate(&UserSession{})
	if err != nil {
		log.Printf("Could not automigrate: [%s]", err.Error())
	}

	err = conn.AutoMigrate(&Dummy{})
	if err != nil {
		log.Printf("Could not automigrate: [%s]", err.Error())
	}

	dummy := Dummy{
		Value: ":)",
	}

	log.Printf("Creating dummy value.")
	conn.Create(&dummy)

	var found Dummy
	log.Printf("Reading dummy value.")
	conn.First(&found)
	log.Printf("Found object: %v", found)

}
