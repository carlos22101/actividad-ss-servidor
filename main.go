package main

import (
	"log"
	"os/exec"
)

func main() {
	log.Println("Iniciando servidores...")

	// Ejecutar el servidor principal
	cmd1 := exec.Command("go", "run", "./servidor_principal/main.go")
	cmd1.Stdout, cmd1.Stderr = log.Writer(), log.Writer()
	go cmd1.Run()

	// Ejecutar el servidor secundario
	cmd2 := exec.Command("go", "run", "./servidor_secundario/main.go")
	cmd2.Stdout, cmd2.Stderr = log.Writer(), log.Writer()
	cmd2.Run()
}