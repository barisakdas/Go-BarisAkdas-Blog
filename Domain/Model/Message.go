package model

import (
	"Go-BarisAkdas-Blog/Infrastructure/Log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Name, Email, Phone, Message, Answer string
	IsDeleted                           bool
	IsReplied                           bool
}

func (message Message) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Message", "Migrate", "Database migration error", err.Error())
	}
	db.AutoMigrate(&message)
}

func (message Message) Add() (Message, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Message", "Add", "Add message error", err.Error())
		return Message{}, err
	}

	db.Create(&message)
	return message, nil
}

func (message Message) Get(where ...interface{}) (Message, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "Get", "Get message error", err.Error())
		return Message{}, err
	}

	db.First(&message, where...)
	return message, nil
}

func (message Message) GetAll(where ...interface{}) ([]Message, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "GetAll", "Get all messages error", err.Error())
		return nil, err
	}

	var messages []Message
	db.Find(&messages, where...)
	return messages, nil
}

func (message Message) Update(data Message) (Message, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "Update", "Update message error", err.Error())
	}

	db.Model(&message).Updates(data)
	return message, err
}

func (message Message) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Article", "Update", "Update message error", err.Error())
		return err
	}

	db.Delete(&message)
	return nil
}
