package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _:= ConnectDB()
	go ShortPolling(db)

	r := gin.Default()
	r.GET("/usuarios", GetUsuariosHandler)
	r.GET("/longpolling", LongPollingUsuarios)
	r.POST("/usuarios", CreateUsuarioHandler)
	r.PUT("/usuarios/:id", UpdateUsuarioHandler)
	r.DELETE("/usuarios/:id", DeleteUsuarioHandler)

	log.Println("Servidor principal corriendo en :8080")
	r.Run(":8080")
}