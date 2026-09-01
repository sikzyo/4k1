package input

import (
	"errors"
	"fmt"
	"strconv"
)

func GetInput() (int, error) {
	var input string
	fmt.Println("Ingresa una respuesta:")
	fmt.Print("-> ")
	fmt.Scan(&input)
	number, err := strconv.Atoi(input)
	if err != nil {
		return -1, errors.New("Solo se aceptan respuestas numéricas")
	}
	return number, nil
}
