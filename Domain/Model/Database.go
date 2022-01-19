package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Dsn string = "host=localhost user=postgres password=Acaba1234!** dbname=blog port=5433 sslmode=disable TimeZone=Asia/Shanghai"

//var Dsn string = "host=db user=postgres password=Acaba1234!** dbname=dnbites port=5433 sslmode=disable TimeZone=Asia/Shanghai"	// For docker-compose change `host`

func CreateDb() {
	db, _ := gorm.Open("postgres", "host=localhost user=postgres password=Acaba1234!** dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai")
	db = db.Exec("CREATE DATABASE blog;")
	if db.Error != nil {
		fmt.Println("Unable to create DB dnbites, attempting to connect assuming it exists...")
	}
}

func Migration() {
	About{}.Migrate()
	Article{}.Migrate()
	Category{}.Migrate()
	Message{}.Migrate()
	Role{}.Migrate()
	User{}.Migrate()
}
