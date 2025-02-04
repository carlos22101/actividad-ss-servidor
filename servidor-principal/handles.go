package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsuariosHandler(c *gin.Context) {
	db := ConnectDB()
	defer db.Close()
	usuarios := GetUsuarios(db)
	c.JSON(http.StatusOK, usuarios)
}

func CreateUsuarioHandler(c *gin.Context) {
	var usuario Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := ConnectDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO usuarios (nombre, contrasena) VALUES (?, ?)", usuario.Nombre, usuario.Contrasena)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}