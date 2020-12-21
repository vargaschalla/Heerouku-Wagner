package models

import "gorm.io/gorm"

type Docente struct {
	gorm.Model
	Codigo string `json:"codigo"`
}
