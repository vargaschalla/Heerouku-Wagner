package models

import "gorm.io/gorm"

type Estudiante struct {
	gorm.Model
	Codigo string `json:"codigo"`
}
