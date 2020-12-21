package models

import "gorm.io/gorm"

type Inscripcion struct {
	gorm.Model
	PlanAcademicoID string        `gorm:"size:191"`
	PlanAcademico   PlanAcademico `gorm :"ForeignKey: Inscripcion"`
	PersonaID       string        `gorm:"size:191"`
	Persona         Persona       `gorm :"ForeignKey: RolPersona"`
}
