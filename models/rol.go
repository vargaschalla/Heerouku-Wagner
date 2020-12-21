package models

import "gorm.io/gorm"

type Rol struct {
	gorm.Model
	Nombre     string `json:"nombre"`
	Estado     string `json:"estado"`
	RolPersona []RolPersona
}
