package input

import (
	"fmt"
	"os"
	"strconv"
)

func GetInput() int {
	var input string
	fmt.Println("Ingresa una respuesta:")
	fmt.Print("->")
	fmt.Scan(&input)
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("-> Algo salio mal al convertir tu respuesta")
		os.Exit(1)
	}
	return number
}
