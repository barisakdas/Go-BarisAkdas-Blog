package model

import (
	"Go-BarisAkdas-Blog/Infrastructure/Log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName, LastName, UserName, Password, BirthDate string
	RoleId                                             int
	IsDeleted                                          bool
}

func (user User) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/User", "Migrate", "Database migration error", err.Error())
	}

	db.AutoMigrate(&user)
}

func (user User) Add() (User, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/User", "Add", "Add user error", err.Error())
		return User{}, err
	}

	db.Create(&user)
	return user, nil
}

func (user User) Get(where ...interface{}) (User, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/User", "Get", "Get user error", err.Error())
		return User{}, err
	}

	db.First(&user, where...)
	return user, nil
}

func (user User) GetAll(where ...interface{}) ([]User, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/User", "GetAll", "Get all users error", err.Error())
		return nil, err
	}

	var users []User
	db.Find(&users, where...)
	return users, nil
}

func (user User) Update(data User) (User, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/User", "Update", "Update user error", err.Error())
		return data, err
	}

	db.Model(&user).Updates(data)
	return user, nil
}

func (user User) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/User", "Update", "Update user error", err.Error())
		return err
	}

	db.Delete(&user)
	return nil
}
