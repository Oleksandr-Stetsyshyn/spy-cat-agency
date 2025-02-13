package models

import (
	"gorm.io/gorm"
)

type Cat struct {
	gorm.Model
	Name              string  `json:"name"`
	YearsOfExperience int     `json:"years_of_experience"`
	Breed             string  `json:"breed"`
	Salary            float64 `json:"salary"`
}
