package models

import "gorm.io/gorm"

type Unidad struct {
	gorm.Model
	Titulo    string `json:"titulo"`
	NumUnidad string `json:"numunidad"`
	Contenido string `json:"contenido"`
	Estado    string `json:"estado"`
	Sesion    []Sesion
}
