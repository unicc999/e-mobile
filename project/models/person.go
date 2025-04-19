package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Surname     string `gorm:"not null" json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Person{})
}
