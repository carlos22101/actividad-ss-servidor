package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsuariosHandler(c *gin.Context) {
	db, _ := ConnectDB()
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
	db, _ := ConnectDB()
	defer db.Close()
	_, err := db.Exec("INSERT INTO usuarios (nombre, contrasena) VALUES (?, ?)", usuario.Nombre, usuario.Contrasena)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
func UpdateUsuarioHandler(c *gin.Context) {
    id := c.Param("id")
    var usuario Usuario
    if err := c.ShouldBindJSON(&usuario); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db, err := ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer db.Close()
    _, err = db.Exec("UPDATE usuarios SET nombre = ?, contrasena = ? WHERE id = ?", usuario.Nombre, usuario.Contrasena, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
}

func DeleteUsuarioHandler(c *gin.Context) {
    id := c.Param("id")
    db, err := ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer db.Close()
    _, err = db.Exec("DELETE FROM usuarios WHERE id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusOK)
}