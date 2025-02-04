package main

import (
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	UsuariosCache []Usuario
	mu            sync.Mutex
	updateChan    = make(chan []Usuario)
)

func GetUsuarios(db *sql.DB) []Usuario {
	rows, _ := db.Query("SELECT id, nombre, contrasena FROM usuarios")
	defer rows.Close()
	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nombre, &u.Contrasena)
		usuarios = append(usuarios, u)
	}
	return usuarios
}

func ShortPolling(db *sql.DB) {
	for {
		usuarios := GetUsuarios(db)
		mu.Lock()
		if !compareUsuarios(UsuariosCache, usuarios) {
			UsuariosCache = usuarios
			updateChan <- usuarios
		}
		mu.Unlock()
		time.Sleep(5 * time.Second)
	}
}

func LongPollingUsuarios(c *gin.Context) {
	nuevosUsuarios := <-updateChan
	c.JSON(http.StatusOK, nuevosUsuarios)
}

func compareUsuarios(oldList, newList []Usuario) bool {
	if len(oldList) != len(newList) {
		return false
	}
	for i := range oldList {
		if oldList[i] != newList[i] {
			return false
		}
	}
	return true
}