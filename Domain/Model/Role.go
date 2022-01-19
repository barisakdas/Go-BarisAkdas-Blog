package model

import (
	"Go-BarisAkdas-Blog/Infrastructure/Log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name, Description                              string
	SelectAuth, CreateAuth, UpdateAuth, DeleteAuth bool
	IsDeleted                                      bool
}

func (role Role) Migrate() {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Role", "Migrate", "Database migration error", err.Error())
	}
	db.AutoMigrate(&role)
}

func (role Role) Add() (Role, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Role", "Add", "Add role error", err.Error())
		return Role{}, err
	}

	db.Create(&role)
	return role, nil
}

func (role Role) Get(where ...interface{}) (Role, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Role", "Get", "Get role error", err.Error())
		return Role{}, err
	}

	db.First(&role, where...)
	return role, nil
}

func (service Role) GetAll(where ...interface{}) ([]Role, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Role", "GetAll", "Get all role error", err.Error())
		return nil, err
	}

	var roles []Role
	db.Find(&roles, where...)
	return roles, nil
}

func (role Role) Update(data Role) (Role, error) {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Role", "Update", "Update role error", err.Error())
		return data, err
	}

	db.Model(&role).Updates(data)
	return role, nil
}

func (role Role) Delete() error {
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		Log.LogJson("Error", "Models/Role", "Update", "Update role error", err.Error())
		return err
	}

	db.Delete(&role)
	return nil
}
