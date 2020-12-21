package models

import "gorm.io/gorm"

type Trabajo struct {
	gorm.Model
	Nombre       string    `json:"nombre"`
	FechaInicio  string    `json:"fechainicio"`
	FechaFin     string    `json:"fechafin"`
	FechaEntrega string    `json:"fechaentrega"`
	Nota         string    `json:"nota"`
	Estado       string    `json:"estado"`
	SecuenciaID  string    `gorm:"size:191"`
	PersonaID    string    `gorm:"size:191"`
	RecursoID    string    `gorm:"size:191"`
	Secuencia    Secuencia `gorm :"ForeignKey: Trabajo"`
	Persona      Persona   `gorm :"ForeignKey: Trabajo"`
	Recurso      Recurso   `gorm :"ForeignKey: Trabajo"`
}
