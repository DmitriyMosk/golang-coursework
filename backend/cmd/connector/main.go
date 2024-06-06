package connector

import (
	"golang-coursework/internal/connector"
	"log"
)

func main() {
	err := connector.Run()
	if err != nil {
		log.Fatalf("Error running connector: %v", err)
	}
}
