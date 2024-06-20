package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Ejercicios")
		fmt.Println("2. Rutina")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la entrada del usuario
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			continue
		}

		// Eliminar los espacios en blanco y convertir a entero
		input = strings.TrimSpace(input)
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entrada inválida, por favor ingrese un número.")
			continue
		}

		// Ejecutar la función correspondiente según la opción seleccionada
		switch option {
		case 1:
			NewListaDeEjercicios().MenuEjercicios()
		case 2:
			NewListaDeRutinas().MenuRutinas()
		case 3:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida, por favor seleccione una opción del 1 al 3.")
		}
	}
}
