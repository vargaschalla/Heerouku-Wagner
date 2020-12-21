package models

import (
	"gorm.io/gorm"
)

type PlanAcademico struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Estado      string `json:"estado"`
	SeccionID   string `gorm:"size:191"`
	PeriodoID   string `gorm:"size:191"`
	CursoID     string `gorm:"size:191"`
	PersonaID   string `gorm:"size:191"`
	Unidad      []Unidad
	Inscripcion []Inscripcion
	Personas    []Persona
	Seccion     Seccion `gorm :"ForeignKey: PlanAcademico"`
	Periodo     Periodo `gorm :"ForeignKey: PlanAcademico"`
	Curso       Curso   `gorm :"ForeignKey: PlanAcademico"`
	Persona     Persona `gorm :"ForeignKey: PlanAcademico"`
}
