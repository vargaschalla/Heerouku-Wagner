package app

import (
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InscripcionIndex(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	lis := []models.Inscripcion{}
	conn.Find(&lis)
	conn.Preload("TipoRecurso").Preload("Sesion").Find(&lis) // Preload("Alumno") carga los objetos Alumno relacionado
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func InscripcionCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	var d models.Inscripcion
	//d := models.Person{Name: c.PostForm("name"), Age: c.PostForm("age")}
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}

func InscripcionGet(c *gin.Context) {

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Inscripcion
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)
}

func InscripcionUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Inscripcion
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

func InscripcionDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Inscripcion

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
