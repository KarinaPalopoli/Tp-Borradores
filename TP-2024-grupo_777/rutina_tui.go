package main

// TUI -> Terminal User Interface

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (lista *ListaDeRutinas) MenuRutinas() {
	for {
		fmt.Println("Menu Rutinas:")
		fmt.Println("1. Agregar Rutinas")
		fmt.Println("2. Borrar Rutinas")
		fmt.Println("3. Consultar Rutinas")
		fmt.Println("4. Modificar Rutinas")
		fmt.Println("5. Listar Rutinas")
		fmt.Println("6. Volver al menú principal")
		fmt.Print("Seleccione una opción: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
			continue
		}
		input = strings.TrimSpace(input)
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Entrada inválida, por favor ingrese un número.")
			continue
		}

		switch option {
		case 1:
			lista.agregarRutina()
		case 2:
			lista.borrarRutina()
		case 3:
			lista.consultarRutina()
		case 4:
			lista.modificarRutina()
		case 5:
			rutinasSlice := make([]*Rutina, 0, len(lista.listaDeRutinas))
			for _, rutina := range lista.listaDeRutinas {
				rutinasSlice = append(rutinasSlice, rutina)
			}
			lista.listarRutinas(rutinasSlice)
		case 6:
			return
		default:
			fmt.Println("Opción no válida, por favor seleccione una opción del 1 al 7.")
		}
	}
}

func (lista *ListaDeRutinas) agregarRutina() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre de la rutina: ")
	nombre, _ := reader.ReadString('\n')
	fmt.Print("Ingrese rutinas separadas por coma: ")
	rutinas, _ := reader.ReadString('\n')

	var ejercicio []*Ejercicio
	for _, nombre := range strings.Split(rutinas, ",") {
		ejer, err := NewListaDeEjercicios().ConsultarEjercicioPorNombre(strings.TrimSpace(nombre))
		if err != nil {
			fmt.Println(err)
			return
		}
		ejercicio = append(ejercicio, ejer)
	}

	err := lista.AgregarRutina(strings.TrimSpace(nombre), ejercicio)
	if err != nil {
		fmt.Println("Error al agregar rutina:", err)
	} else {
		fmt.Println("Rutina agregado exitosamente.")
		lista.guardarRutinasEnCSV()
	}
}

func (lista *ListaDeRutinas) borrarRutina() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre del rutina a borrar: ")
	nombre, _ := reader.ReadString('\n')
	err := lista.BorrarRutina(strings.TrimSpace(nombre))
	if err != nil {
		fmt.Println("Error al borrar rutina:", err)
	} else {
		fmt.Println("Rutina borrado exitosamente.")
		lista.guardarRutinasEnCSV()
	}
}

// Función para listar todos los rutinas
func (lista *ListaDeRutinas) listarRutinas(listaDeRutinas []*Rutina) {
	fmt.Println("Lista de Rutinas:")

	for _, rutina := range listaDeRutinas {
		var ejerciciosTotales []string
		for _, nombre := range  rutina.ejerciciosTotales{
			ejerciciosTotales = append(ejerciciosTotales, nombre.nombre)
		}

		fmt.Printf("Nombre: %s\n", rutina.nombre)
		fmt.Printf("Duracion: %d\n", rutina.duracion)
		fmt.Printf("ejercicios Totales: %v \n", strings.Join(ejerciciosTotales,","),)
		fmt.Printf("tipoDeEjercicios: %v\n", rutina.tipoDeEjercicios)
		fmt.Printf("caloriasQuemadasTotales: %v\n", rutina.caloriasQuemadasTotales)
		fmt.Printf("dificultad: %v\n", rutina.dificultad)
		fmt.Println("-----------------------")
	}
}

func (lista *ListaDeRutinas) consultarRutina() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre del rutina a consultar: ")
	nombre, _ := reader.ReadString('\n')
	rutina, err := lista.ConsultarRutina(strings.TrimSpace(nombre))
	if err != nil {
		fmt.Println("Error al consultar rutina:", err)
	} else {
		fmt.Printf("Rutina: %+v\n", rutina)
	}
}

func (lista *ListaDeRutinas) modificarRutina() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre de la rutina a modificar: ")
	nombre, _ := reader.ReadString('\n')
	fmt.Print("Ingrese rutinas separadas por coma: ")
	rutinas, _ := reader.ReadString('\n')

	var ejercicio []*Ejercicio
	for _, nombre := range strings.Split(rutinas, ",") {
		ejer, err := NewListaDeEjercicios().ConsultarEjercicioPorNombre(strings.TrimSpace(nombre))
		if err != nil {
			fmt.Println(err)
			return
		}
		ejercicio = append(ejercicio, ejer)
	}

	err := lista.ModificarRutina(strings.TrimSpace(nombre), ejercicio)
	if err != nil {
		fmt.Println("Error al modifico la rutina:", err)
	} else {
		fmt.Println("Rutina modifico exitosamente.")
		lista.guardarRutinasEnCSV()
	}
}
