package app

import (
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RolPersonaLista(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	lis := []models.RolPersona{}
	conn.Find(&lis)
	conn.Preload("Persona").Preload("Rol").Find(&lis) // Preload("Alumno") carga los objetos Alumno relacionado
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})
}
func RolPersonaGet(c *gin.Context) {

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.RolPersona
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)
}
func RolPersonaCreate(c *gin.Context) {
	var p models.RolPersona
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	conn.Create(&p)
	c.JSON(http.StatusOK, &p)
}

func RolPersonaUpdate(c *gin.Context) {
	var d models.RolPersona
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}

func RolPersonaDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.RolPersona

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
