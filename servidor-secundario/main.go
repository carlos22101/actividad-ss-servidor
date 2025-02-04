package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Servidor secundario corriendo en :8081")
	go StartPolling()
	select {}
}