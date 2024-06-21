package main

// TUI -> Terminal User Interface

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func (lista *ListaDeEjercicios) MenuEjercicios() {
	for {
		fmt.Println("Menu Ejercicios:")
		fmt.Println("1. Agregar Ejercicio")
		fmt.Println("2. Borrar Ejercicio")
		fmt.Println("3. Consultar Ejercicio")
		fmt.Println("4. Filtrar Ejercicios")
		fmt.Println("5. Modificar Ejercicio")
		fmt.Println("6. Listar Ejercicios")
		fmt.Println("7. Volver al menú principal")
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
			lista.agregarEjercicio()
		case 2:
			lista.borrarEjercicio()
		case 3:
			lista.consultarEjercicio()
		case 4:
			lista.filtrarEjercicios()
		case 5:
			lista.modificarEjercicio()
		case 6:
			ejerciciosSlice := make([]*Ejercicio, 0, len(lista.listaDeEjercicios))
			for _, ejercicio := range lista.listaDeEjercicios {
				ejerciciosSlice = append(ejerciciosSlice, ejercicio)
			}
			lista.listarEjercicios(ejerciciosSlice)
		case 7:
			return
		default:
			fmt.Println("Opción no válida, por favor seleccione una opción del 1 al 7.")
		}
	}
}

func (lista *ListaDeEjercicios) agregarEjercicio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre del ejercicio: ")
	nombre, _ := reader.ReadString('\n')
	fmt.Print("Ingrese descripcion: ")
	descripcion, _ := reader.ReadString('\n')
	fmt.Print("Ingrese tiempo (minutos): ")
	tiempoStr, _ := reader.ReadString('\n')
	fmt.Print("Ingrese calorias: ")
	caloriasStr, _ := reader.ReadString('\n')
	fmt.Print("Ingrese tipos de ejercicio (separados por coma): ")
	tiposStr, _ := reader.ReadString('\n')
	fmt.Print("Ingrese puntos por tipo de ejercicio (separados por coma): ")
	puntosStr, _ := reader.ReadString('\n')
	fmt.Print("Ingrese dificultad: ")
	dificultad, _ := reader.ReadString('\n')

	tiempo, _ := strconv.Atoi(strings.TrimSpace(tiempoStr))
	calorias, _ := strconv.Atoi(strings.TrimSpace(caloriasStr))
	tipos := strings.Split(strings.TrimSpace(tiposStr), ",")
	puntosStrs := strings.Split(strings.TrimSpace(puntosStr), ",")
	puntos := make([]int, len(puntosStrs))
	for i, p := range puntosStrs {
		puntos[i], _ = strconv.Atoi(strings.TrimSpace(p))
	}

	err := lista.AgregarEjercicio(strings.TrimSpace(nombre), strings.TrimSpace(descripcion), tiempo, calorias, tipos, puntos, strings.TrimSpace(dificultad))
	if err != nil {
		fmt.Println("Error al agregar ejercicio:", err)
	} else {
		fmt.Println("Ejercicio agregado exitosamente.")
		lista.guardarEjerciciosEnCSV()

	}
}

func (lista *ListaDeEjercicios) borrarEjercicio() {
	// 	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Presiona Enter para ver más nombres de ejercicios (o q + Enter para salir)")

	// Crear un scanner para leer la entrada estándar

	fmt.Println("Lista de Ejercicios:")
	for _, ejercicio := range lista.listaDeEjercicios {
		fmt.Fprintf(os.Stdout, "Nombre: %s\n", []any{ejercicio.nombre}...)

		fmt.Println("-----------------------")
	}

	fmt.Println("¿Qué ejercicio deseas eliminar? (escribe el nombre del ejercicio)")

	// Limpiar la pantalla antes de mostrar más nombres
	fmt.Print("\033[H\033[2J")

	fmt.Println("Fin del listado de nombres de ejercicios.")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre del ejercicio a borrar: ")

	nombre, _ := reader.ReadString('\n')
	err := lista.BorrarEjercicio(strings.TrimSpace(nombre))
	if err != nil {
		fmt.Println("Error al borrar ejercicio:", err)
	} else {
		fmt.Println("Ejercicio borrado exitosamente.")
		lista.guardarEjerciciosEnCSV()
	}
}

// Función para listar los ejercicios de a dos por vez
func (lista *ListaDeEjercicios) listarEjercicios(listaDeEjercicios []*Ejercicio) {
	fmt.Println("Presiona Enter para ver más ejercicios (o q + Enter para salir)")

	// Crear un scanner para leer la entrada estándar
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < len(listaDeEjercicios); i += 2 {
		ejercicio1 := listaDeEjercicios[i]
		fmt.Printf("Nombre: %s\n", ejercicio1.nombre)
		fmt.Printf("Descripción: %s\n", ejercicio1.descripcion)
		fmt.Printf("Tiempo: %d minutos\n", ejercicio1.tiempo)
		fmt.Printf("Calorías: %d\n", ejercicio1.calorias)
		fmt.Printf("Tipo de Ejercicio: %v\n", ejercicio1.tipoDeEjercicio)
		fmt.Printf("Puntos por Tipo de Ejercicio: %v\n", ejercicio1.puntosPorTipoDeEjercicio)
		fmt.Printf("Dificultad: %s\n", ejercicio1.dificultad)
		fmt.Println("-----------------------")

		// Verificar si hay un segundo ejercicio para mostrar
		if i+1 < len(listaDeEjercicios) {
			ejercicio2 := listaDeEjercicios[i+1]
			fmt.Printf("Nombre: %s\n", ejercicio2.nombre)
			fmt.Printf("Descripción: %s\n", ejercicio2.descripcion)
			fmt.Printf("Tiempo: %d minutos\n", ejercicio2.tiempo)
			fmt.Printf("Calorías: %d\n", ejercicio2.calorias)
			fmt.Printf("Tipo de Ejercicio: %v\n", ejercicio2.tipoDeEjercicio)
			fmt.Printf("Puntos por Tipo de Ejercicio: %v\n", ejercicio2.puntosPorTipoDeEjercicio)
			fmt.Printf("Dificultad: %s\n", ejercicio2.dificultad)
			fmt.Println("-----------------------")
		}

		// Esperar a que el usuario presione Enter para mostrar más ejercicios
		fmt.Print("Presiona Enter para ver más ejercicios...")
		scanner.Scan()
		if scanner.Text() == "q" {
			break // Salir del bucle si el usuario presiona 'q'
		}

		// Limpiar la pantalla antes de mostrar más ejercicios
		clearScreen()
	}

	fmt.Println("Fin del listado de ejercicios.")
}

// Función para limpiar la pantalla
func clearScreen() {
	cmd := exec.Command("clear") // Linux/MacOS
	if err := cmd.Run(); err != nil {
		fmt.Println("No se pudo limpiar la pantalla:", err)
	}
}

// Función para listar todos los ejercicios

// func (lista *ListaDeEjercicios) listarEjercicios(listaDeEjercicios []*Ejercicio) {
// 	fmt.Println("Lista de Ejercicios:")
// 	for _, ejercicio := range listaDeEjercicios {
// 		fmt.Printf("Nombre: %s\n", ejercicio.nombre)
// 		fmt.Printf("Descripción: %s\n", ejercicio.descripcion)
// 		fmt.Printf("Tiempo: %d minutos\n", ejercicio.tiempo)
// 		fmt.Printf("Calorías: %d\n", ejercicio.calorias)
// 		fmt.Printf("Tipo de Ejercicio: %v\n", ejercicio.tipoDeEjercicio)
// 		fmt.Printf("Puntos por Tipo de Ejercicio: %v\n", ejercicio.puntosPorTipoDeEjercicio)
// 		fmt.Printf("Dificultad: %s\n", ejercicio.dificultad)
// 		fmt.Println("-----------------------")
// 	}
// }

func (lista *ListaDeEjercicios) consultarEjercicio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre del ejercicio a consultar: ")
	nombre, _ := reader.ReadString('\n')
	ejercicio, err := lista.ConsultarEjercicioPorNombre(strings.TrimSpace(nombre))
	if err != nil {
		fmt.Println("Error al consultar ejercicio:", err)
	} else {
		fmt.Printf("Ejercicio: %+v\n", ejercicio)
	}
}

func (lista *ListaDeEjercicios) filtrarEjercicios() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese tipo de ejercicio (dejar vacío para ignorar): ")
	tipo, _ := reader.ReadString('\n')
	fmt.Print("Ingrese dificultad (dejar vacío para ignorar): ")
	dificultad, _ := reader.ReadString('\n')
	fmt.Print("Ingrese calorías mínimas (0 para ignorar): ")
	minCaloriasStr, _ := reader.ReadString('\n')
	minCalorias, _ := strconv.Atoi(strings.TrimSpace(minCaloriasStr))

	ejercicios, err := lista.FiltrarEjercicios(strings.TrimSpace(tipo), strings.TrimSpace(dificultad), minCalorias)
	if err != nil {
		fmt.Println("Error al filtrar ejercicios:", err)
	} else {
		lista.listarEjercicios(ejercicios)
	}
}

func (lista *ListaDeEjercicios) modificarEjercicio() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese nombre del ejercicio a modificar: ")
	nombre, _ := reader.ReadString('\n')
	fmt.Print("Ingrese nueva descripcion: ")
	nuevaDescripcion, _ := reader.ReadString('\n')
	fmt.Print("Ingrese nuevo tiempo (minutos): ")
	nuevoTiempoStr, _ := reader.ReadString('\n')
	fmt.Print("Ingrese nuevas calorias (por minuto de éste ejercicio): ")
	nuevasCaloriasStr, _ := reader.ReadString('\n')
	fmt.Print("Ingrese nuevos tipos de ejercicio (separados por coma): ")
	nuevosTiposStr, _ := reader.ReadString('\n')
	fmt.Print("Ingrese nuevos puntos por tipo de ejercicio (separados por coma): ")
	nuevosPuntosStr, _ := reader.ReadString('\n')
	fmt.Print("Ingrese  dificultad: principiante / intermedio / avanzado (SÓLO UNA PUEDE ELEGIR): ")
	nuevaDificultad, _ := reader.ReadString('\n')

	nuevoTiempo, _ := strconv.Atoi(strings.TrimSpace(nuevoTiempoStr))
	nuevasCalorias, _ := strconv.Atoi(strings.TrimSpace(nuevasCaloriasStr))
	nuevosTipos := strings.Split(strings.TrimSpace(nuevosTiposStr), ",")
	nuevosPuntosStrs := strings.Split(strings.TrimSpace(nuevosPuntosStr), ",")
	nuevosPuntos := make([]int, len(nuevosPuntosStrs))
	for i, p := range nuevosPuntosStrs {
		nuevosPuntos[i], _ = strconv.Atoi(strings.TrimSpace(p))
	}

	err := lista.ModificarEjercicio(strings.TrimSpace(nombre), strings.TrimSpace(nuevaDescripcion), nuevoTiempo, nuevasCalorias, nuevosTipos, nuevosPuntos, strings.TrimSpace(nuevaDificultad))
	if err != nil {
		fmt.Println("Error al modificar ejercicio:", err)
	} else {
		fmt.Println("Ejercicio modificado exitosamente.")
		lista.guardarEjerciciosEnCSV()
	}
}
