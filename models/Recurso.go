package models

import "gorm.io/gorm"

type Recurso struct {
	gorm.Model
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	Estado        string `json:"estado"`
	TipoRecursoID string `gorm:"size:191"`
	SesionID      string `gorm:"size:191"`
	Trabajo       []Trabajo
	TipoRecurso   TipoRecurso `gorm :"ForeignKey: Recurso"`
	Sesion        Sesion      `gorm :"ForeignKey: Recurso"`
}
