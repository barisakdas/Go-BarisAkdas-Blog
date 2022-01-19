package model

import (
	"Go-BarisAkdas-Blog/Infrastructure/Log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name      string
	IsDeleted bool
}

func (category Category) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Category", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&category)
}

func (category Category) Add() (Category, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Category", "Add", "Add category error", err.Error())
		return Category{}, err
	}

	db.Create(&category)
	return category, nil
}

func (category Category) Get(where ...interface{}) (Category, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Category", "Get", "Get category error", err.Error())
		return category, err
	}

	db.First(&category, where...)
	return category, nil
}

func (category Category) GetAll(where ...interface{}) ([]Category, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Category", "GetAll", "Get all categories error", err.Error())
		return nil, err
	}

	var categories []Category
	db.Find(&categories, where...)
	return categories, nil
}

func (category Category) Update(data Category) (Category, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Category", "Update", "Update category error", err.Error())
		return data, err
	}

	db.Model(&category).Updates(data)
	return data, nil
}

func (category Category) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Category", "Update", "Update category error", err.Error())
		return err
	}

	db.Delete(&category)
	return nil
}
