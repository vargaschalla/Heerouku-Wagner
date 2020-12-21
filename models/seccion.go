package models

import "gorm.io/gorm"

type Seccion struct {
	gorm.Model
	Alias         string `json:"alias"`
	Nombre        string `json:"nombre"`
	Estado        string `json:"estado"`
	PlanAcademico []PlanAcademico
}
