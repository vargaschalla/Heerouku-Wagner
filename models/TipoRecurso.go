package models

import "gorm.io/gorm"

type TipoRecurso struct {
	gorm.Model
	Nombre  string `json:"nombre"`
	Estado  string `json:"estado"`
	Recurso []Recurso
}
