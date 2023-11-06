package cmd

import (
	"awesomeProject/internal"
	"awesomeProject/model"
	"awesomeProject/pkg"
	"gorm.io/gorm"
	"log"
)

func InitializeDB(connOptional ...*gorm.DB) {
	var conn *gorm.DB
	var err error
	if len(connOptional) == 0 {
		conn, err = internal.Connection()
		if err != nil {
			log.Fatal("Couldn't initialize db")
		}
	} else {
		conn = connOptional[0]
	}

	settings := []model.Setting{
		{SettingName: "first_access", SettingValue: "no"},
		{SettingName: "template", SettingValue: "default"},
		{SettingName: "api_type", SettingValue: pkg.NO_API}, // Accepted values: no_api, with_api and only_api, see pkg/api-type.go
	}
	for _, setting := range settings {
		conn.Create(&setting)
	}
}
