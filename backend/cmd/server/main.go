package server

import (
	"golang-coursework/internal/server"
	"log"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
