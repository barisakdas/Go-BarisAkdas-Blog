package model

import (
	"Go-BarisAkdas-Blog/Infrastructure/Log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type About struct {
	gorm.Model
	Title, Description string
	IsDeleted          bool
}

func (about About) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/About", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&about)
}

func (about About) GetAll(where ...interface{}) ([]About, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/About", "GetAll", "Get all about error", err.Error())
		return nil, err
	}

	var abouts []About
	db.Find(&abouts, where...)
	return abouts, nil
}

func (about About) Get(where ...interface{}) (About, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/About", "Get", "Get about error", err.Error())
		return about, err
	}

	db.First(&about, where...)
	return about, nil
}

func (about About) Add() (About, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/About", "Add", "Add about error", err.Error())
		return About{}, err
	}

	db.Create(&about)
	return about, nil
}

func (about About) Update(data About) (About, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/About", "Update", "Update about error", err.Error())
		return data, err
	}

	db.Model(&about).Updates(data)
	return data, nil
}

func (about About) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/About", "Delete", "Delete about error", err.Error())
		return err
	}

	db.Delete(&about)
	return nil
}
