package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Person struct {
	ID          int    `gorm:"primaryKey;autoIncrement" json:"id"`
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

func (p Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Surname     string `json:"surname"`
		Patronymic  string `json:"patronymic"`
		Age         int    `json:"age"`
		Gender      string `json:"gender"`
		Nationality string `json:"nationality"`
	}{
		ID:          p.ID,
		Name:        p.Name,
		Surname:     p.Surname,
		Patronymic:  p.Patronymic,
		Age:         p.Age,
		Gender:      p.Gender,
		Nationality: p.Nationality,
	})
}
