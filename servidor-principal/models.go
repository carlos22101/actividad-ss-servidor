package main

type Usuario struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Contrasena string `json:"contrasena"`
}