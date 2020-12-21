package models

import "gorm.io/gorm"

type Secuencia struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Estado      string `json:"estado"`
	Abrebiatura string `json:"abrebiatura"`
	Trabajo     []Trabajo
}
