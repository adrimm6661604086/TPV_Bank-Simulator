package utils

import (
	"log"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func JsonEncoder(message string) Response {
	return Response{
		Message: message,
		Status:  200,
	}
}

func WithLogPrefix(prefix string, fn func() interface{}) interface{} {
	originalPrefix := log.Prefix()
	log.SetPrefix(prefix)
	defer log.SetPrefix(originalPrefix)
	return fn()
}
