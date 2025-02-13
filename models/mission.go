package models

import (
	"gorm.io/gorm"
)

type Mission struct {
	gorm.Model
	CatID    uint     `json:"cat_id"`
	Targets  []Target `json:"targets" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Complete bool     `json:"complete"`
}
