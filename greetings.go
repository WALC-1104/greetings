package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Nombre Vacío")
	}

	//Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randon format devuelve uno de varios formatos de mensajes de saludo
// se selecciona de forma aleatoria
func randomFormat() string {
	formats := []string{
		"Hola, %v! Bienvenido!",
		"Que bueno verte, %v!",
		"Saludo, %v! encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
