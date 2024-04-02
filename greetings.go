package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("nombre vacio")
	}
	//Devuelve la cadena formateadas
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
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
func randomFormat() string {
	formats := []string{
		"Hola %v",
		"Aloha %v",
		"Bienvenid@, Sr/Sra %v"}
	return formats[rand.Intn(len(formats))]
}
