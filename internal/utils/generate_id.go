package utils

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length   = 32
)

func GenerateId() string {
	id, err := gonanoid.Generate(alphabet, length)
	if err != nil {
		log.Fatal("Error generating id: ", err)
	}

	return id
}
