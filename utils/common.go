package utils

import "log"

func HandleError(message string, filename string, err error) {
	if err != nil {
		log.Fatalf("%s [%s]: %v", message, filename, err)
	}
}
