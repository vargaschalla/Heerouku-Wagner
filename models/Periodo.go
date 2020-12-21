package models

import (
	"gorm.io/gorm"
)

type Periodo struct {
	gorm.Model
	Nombre        string `json:"nombre"`
	Estado        string `json:"estado"`
	PlanAcademico []PlanAcademico
}
