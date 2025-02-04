package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Usuario struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Contrasena string `json:"contrasena"`
}

var cacheUsuarios []Usuario

func StartPolling() {
	for {
		usuarios := FetchUsuarios()
		if !compareUsuarios(cacheUsuarios, usuarios) {
			cacheUsuarios = usuarios
			log.Println("Cambio detectado, solicitando long polling...")
			usuariosActualizados := FetchLongPolling()
			cacheUsuarios = usuariosActualizados
		}
		time.Sleep(5 * time.Second)
	}
}

func FetchUsuarios() []Usuario {
	resp, err := http.Get("http://localhost:8080/usuarios")
	if err != nil {
		log.Println("Error obteniendo usuarios:", err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var usuarios []Usuario
	json.Unmarshal(body, &usuarios)
	return usuarios
}

func FetchLongPolling() []Usuario {
	resp, err := http.Get("http://localhost:8080/longpolling")
	if err != nil {
		log.Println("Error en long polling:", err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var usuarios []Usuario
	json.Unmarshal(body, &usuarios)
	log.Println("Usuarios actualizados recibidos:", usuarios)
	return usuarios
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