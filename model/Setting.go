package model

type Setting struct {
	ID           uint
	SettingName  string
	SettingValue string `gorm:"comment:For false values use no;"`
}
