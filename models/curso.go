package models

import "gorm.io/gorm"

type Curso struct {
	gorm.Model
	Nombre        string `json:"nombre"`
	Estado        string `json:"estado"`
	PlanAcademico []PlanAcademico
}
