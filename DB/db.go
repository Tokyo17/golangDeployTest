package db

import (
	"fmt"
	"log"

	"belajar-lagi/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	host := "ec2-44-198-196-149.compute-1.amazonaws.com"
	dbName := "d8q243p4pns841"
	port := "5432"
	user := "oxbfvmecwdpvlh"
	pass := "41c65db8db5facfa90931eeb4d6dec1c7b450d87bbac2e4bfda89338e4b860f9"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta",
		host, user, pass, dbName, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	// db.AutoMigrate(&Coba{})

	// cek := db.Migrator().HasTable(&Coba{})
	// fmt.Println(cek)
	db.Migrator().CreateTable(&model.Coba{})
	// coba := model.Coba{}
	// db.Find(&coba)
	// fmt.Println(coba)
	return db
}
