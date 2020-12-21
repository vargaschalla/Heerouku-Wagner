package app

import (
	"PROYECTintegrador/ProyectoGOI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SeccionIndex(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	lis := []models.Seccion{}
	conn.Find(&lis)
	conn.Preload("Seccion").Preload("Periodo").Preload("Curso").Preload("Persona").Find(&lis) // Preload("Alumno") carga los objetos Alumno relacionado
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func SeccionCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	var d models.Seccion
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

func SeccionGet(c *gin.Context) {

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Seccion
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)
}

func SeccionUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Seccion
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

func SeccionDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Seccion

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
