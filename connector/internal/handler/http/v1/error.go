package v1

import (
	"fmt"
	"golang-coursework/connector/pkg/logger"
	"net/http"
)

func errorWriter(w http.ResponseWriter, log *logger.Logger, message string, status int) {
	fmt.Println(message)
	log.Log(logger.ERROR, message)
	w.WriteHeader(status)
}
