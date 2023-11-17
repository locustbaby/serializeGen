package utils

import "log"

func HandleError(context string, err error, message string) {
	if err != nil {
		log.Fatalf("%s [%s]: %v", message, context, err)
	}
}
