package model

import (
	"Go-BarisAkdas-Blog/Infrastructure/Log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title, Slug, Description, Content, PictureUrl string
	CategoryID                                    int
	IsDeleted                                     bool
}

func (article Article) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&article)
}

func (article Article) Add() (Article, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "Add", "Add article error", err.Error())
		return Article{}, err

	}

	db.Create(&article)
	return article, nil
}

func (article Article) Get(where ...interface{}) (Article, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "Get", "Get article error", err.Error())
		return article, err
	}

	db.First(&article, where...)
	return article, nil
}

func (article Article) GetAll(where ...interface{}) ([]Article, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "GetAll", "Get all articles error", err.Error())
		return nil, err
	}

	var articles []Article
	db.Find(&articles, where...)
	return articles, nil
}

func (article Article) Update(data Article) (Article, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
		return data, err
	}

	db.Model(&article).Updates(data)
	return data, nil
}

func (article Article) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "Update", "Update article error", err.Error())
		return err
	}

	db.Delete(&article)
	return nil
}
