package app

import (
	"net/http"

	"PROYECTintegrador/ProyectoGOI/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Declaring layout constant
const layout = "2006-Jan-02"

//CRUD for items table
func PlanAcademicoIndex(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	lis := []models.PlanAcademico{}
	conn.Find(&lis)
	conn.Preload("Seccion").Preload("Periodo").Preload("Curso").Preload("Persona").Find(&lis) // Preload("Alumno") carga los objetos Alumno relacionado
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func PlanAcademicoCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	var d models.PlanAcademico
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}

func PlanAcademicoGet(c *gin.Context) {

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.PlanAcademico
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)
}

func PlanAcademicoUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.PlanAcademico
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

func PlanAcademicoDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.PlanAcademico

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
