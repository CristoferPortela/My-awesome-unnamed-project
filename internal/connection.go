package internal

import (
	"awesomeProject/model"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// TODO remove this later
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func Connection() (*gorm.DB, error) {

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(conn)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database error")
		return nil, err
	}
	if err := db.AutoMigrate(&model.User{}, &model.Setting{}); err != nil {
		log.Fatal("Auto migration error")
	}

	return db, nil
}
