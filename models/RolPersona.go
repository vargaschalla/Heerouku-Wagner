package models

import "gorm.io/gorm"

type RolPersona struct {
	gorm.Model
	PersonaID string  `gorm:"size:191"`
	RolID     string  `gorm:"size:191"`
	Persona   Persona `gorm :"ForeignKey: RolPersona"`
	Rol       Rol     `gorm :"ForeignKey: RolPersona"`
}
