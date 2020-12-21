package models

import "gorm.io/gorm"

type Sesion struct {
	gorm.Model
	Titulo   string `json:"titulo"`
	Tema     string `json:"tema"`
	Estado   string `json:"estado"`
	Recurso  []Recurso
	UnidadID string `gorm:"size:191"`
	Unidad   Unidad `gorm :"ForeignKey: Sesion"`
}
