package common

import (
	"errors"
	"log"
)

var (
	RecordNotFound = errors.New("record not founding")
)
func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}