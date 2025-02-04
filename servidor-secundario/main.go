package main

import (
	"log"
	
)

func main() {
	log.Println("Servidor secundario corriendo en :8081")
	go StartPolling()
	select {}
}