package models

import "gorm.io/gorm"

type Target struct {
	gorm.Model
	Name      string `json:"name"`
	Country   string `json:"country"`
	Notes     string `json:"notes"`
	Complete  bool   `json:"complete"`
	MissionID uint   `json:"mission_id"`
}
