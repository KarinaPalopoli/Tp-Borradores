package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

// Estructura de Rutina
type Rutina struct {
	nombre                  string
	duracion                int
	ejerciciosTotales       []*Ejercicio
	tipoDeEjercicios        string // fuerza, cardio, balance
	caloriasQuemadasTotales int
	dificultad              string // principiante, intermedio, avanzado
}

type RutinaCSV struct {
	Nombre                  string `csv:"nombre"`
	Duracion                int    `csv:"duracion"`
	EjerciciosTotales       string `csv:"ejerciciosTotales"`
	TipoDeEjercicios        string `csv:"tipoDeEjercicios"`
	CaloriasQuemadasTotales int    `csv:"caloriasQuemadasTotales"`
	Dificultad              string `csv:"dificultad"`
}

// Función auxiliar para calcular la duración total de una rutina
func calcularDuracion(ejercicios []*Ejercicio) int {
	duracion := 0
	for _, ejercicio := range ejercicios {
		duracion += ejercicio.tiempo
	}
	return duracion
}

// Función auxiliar para calcular las calorías quemadas totales de una rutina
func calcularCaloriasTotales(ejercicios []*Ejercicio) int {
	calorias := 0
	for _, ejercicio := range ejercicios {
		calorias += ejercicio.calorias
	}
	return calorias
}

// Función auxiliar para calcular el tipo de ejercicios más frecuentes
func calcularTipoEjercicios(ejercicios []*Ejercicio) string {
	// Usamos un map para registrar la cantidad de veces que aparece cada tipo de ejercicio
	frecuenciaTipoEjercicio := make(map[string]int)
	// Registrar la cantidad de veces que aparece cada tipo de ejercicio
	for _, ejercicio := range ejercicios {
		for _, tipo := range ejercicio.tipoDeEjercicio {
			frecuenciaTipoEjercicio[tipo]++
		}
	}
	// Encontrar el tipo de ejercicio más frecuente
	frecuenciaMaxima := 0
	tipoMasFrecuente := ""
	for tipo, frecuencia := range frecuenciaTipoEjercicio {
		if frecuencia > frecuenciaMaxima {
			frecuenciaMaxima = frecuencia
			tipoMasFrecuente = tipo
		}
	}
	return tipoMasFrecuente
}

// Función auxiliar para calcular la dificultad más frecuentes
func calcularDificultadEjercicios(ejercicios []*Ejercicio) string {
	// Usamos un map para registrar la cantidad de veces que aparece cada dificultad
	dificultades := make(map[string]int)
	// Registrar la cantidad de veces que aparece cada dificultad
	for _, ejercicio := range ejercicios {
		dificultades[string(ejercicio.dificultad)]++
	}
	// Encontrar la dificultad más frecuente
	frecuenciaMaxima := 0
	dificultadMasFrecuente := ""
	for dificultad, frecuencia := range dificultades {
		if frecuencia > frecuenciaMaxima {
			frecuenciaMaxima = frecuencia
			dificultadMasFrecuente = dificultad
		}
	}
	return dificultadMasFrecuente
}

// Estructura para almacenar las rutinas
type ListaDeRutinas struct {
	listaDeRutinas map[string]*Rutina
}

// Inicializa una ListaDeRutinas y crea el map vacío
func NewListaDeRutinas() *ListaDeRutinas {
	listaTemp := &ListaDeRutinas{listaDeRutinas: make(map[string]*Rutina)}

	err := listaTemp.cargarEjerciciosDesdeCSV()
	if err != nil {
		log.Fatal("Error: ", err)

	}

	return listaTemp

}

func (lista *ListaDeRutinas) cargarEjerciciosDesdeCSV() error {
	// Abrir el archivo CSV para lectura
	file, err := os.OpenFile("rutina.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// Crear un slice para almacenar las rutinas que se van a cargar
	var rutinasCsv []*RutinaCSV

	// Leer las rutinas desde el archivo CSV
	if err := gocsv.UnmarshalFile(file, &rutinasCsv); err != nil {
		return err
	}

	// Agregar las rutinas cargadas al mapa listaDeRutinas
	for _, rutina := range rutinasCsv {
		var ejerciciosTotales []*Ejercicio

		// Dividir la cadena de ejercicios totales por comas
		arrayNombre := strings.Split(rutina.EjerciciosTotales, ",")

		// Iterar sobre los ejercicios usando un for range
		for _, ejercicio := range arrayNombre {
			ObjetoEjercicios, err := NewListaDeEjercicios().ConsultarEjercicioPorNombre(ejercicio)
			if err != nil {
				return err
			}
			ejerciciosTotales = append(ejerciciosTotales, ObjetoEjercicios)
		}

		lista.listaDeRutinas[rutina.Nombre] = &Rutina{
			nombre:                  rutina.Nombre,
			duracion:                rutina.Duracion,
			ejerciciosTotales:       ejerciciosTotales,
			tipoDeEjercicios:        rutina.TipoDeEjercicios,
			caloriasQuemadasTotales: rutina.CaloriasQuemadasTotales,
			dificultad:              rutina.Dificultad,
		}
	}

	fmt.Printf("\n\n--------------\n\n Se cargaron %d registros del CSV \n\n-----------------\n\n", len(lista.listaDeRutinas))

	return nil
}

func (lista *ListaDeRutinas) guardarRutinasEnCSV() error {
	// Crear un slice para almacenar las rutinas
	var rutinasCsv  []*RutinaCSV
	for _, rutina := range lista.listaDeRutinas {

		var ejerciciosTotales []string
		for _, nombre := range  rutina.ejerciciosTotales{
			ejerciciosTotales = append(ejerciciosTotales, nombre.nombre)
		}

		rut := &RutinaCSV{
			Nombre:                  rutina.nombre,
			Duracion:                rutina.duracion,
			EjerciciosTotales:       strings.Join(ejerciciosTotales,","),
			TipoDeEjercicios:        rutina.tipoDeEjercicios,
			CaloriasQuemadasTotales: rutina.caloriasQuemadasTotales,
			Dificultad:              rutina.dificultad,
		}


		rutinasCsv = append(rutinasCsv, rut)
	}

	// Abrir el archivo CSV para escritura
	file, err := os.OpenFile("rutina.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// Escribir las rutinas al archivo CSV
	err = gocsv.MarshalFile(&rutinasCsv, file)
	if err != nil {
		return err
	}

	return nil
}

// AgregarRutina crea una rutina y la agrega al map de listaDeRutinas
func (lista *ListaDeRutinas) AgregarRutina(nombre string, ejerciciosTotales []*Ejercicio) error {
	// Validar que la rutina no exista previamente
	if _, existe := lista.listaDeRutinas[nombre]; existe {
		return errors.New("la rutina ya existe")
	}
	if len(ejerciciosTotales) == 0 {
		return errors.New("una rutina debe contener al menos 1 ejercicio")
	}
	duracionRutina := calcularDuracion(ejerciciosTotales)
	caloriasRutina := calcularCaloriasTotales(ejerciciosTotales)
	tipoEjerciciosRutina := calcularTipoEjercicios(ejerciciosTotales)
	dificultadRutina := calcularDificultadEjercicios(ejerciciosTotales)
	rutina := &Rutina{
		nombre:                  nombre,
		duracion:                duracionRutina,
		ejerciciosTotales:       ejerciciosTotales,
		tipoDeEjercicios:        tipoEjerciciosRutina,
		caloriasQuemadasTotales: caloriasRutina,
		dificultad:              dificultadRutina,
	}
	lista.listaDeRutinas[nombre] = rutina
	return nil
}

// BorrarRutina elimina el par key value, a partir de la key indicada
func (lista *ListaDeRutinas) BorrarRutina(nombre string) error {
	// Validar que la rutina exista
	if _, existe := lista.listaDeRutinas[nombre]; !existe {
		return errors.New("la rutina no existe")
	}
	delete(lista.listaDeRutinas, nombre)
	return nil
}

// ConsultarRutina busca la rutina a partir de la key indicada y devuelve la Rutina
func (lista *ListaDeRutinas) ConsultarRutina(nombre string) (*Rutina, error) {
	// Validar que la rutina exista
	rutina, existe := lista.listaDeRutinas[nombre]
	if !existe {
		return nil, errors.New("la rutina no existe")
	}
	return rutina, nil
}

// ModificarRutina permite modificar los valores de una rutina,
// a partir de identificar la misma a partir de la key indicada
func (lista *ListaDeRutinas) ModificarRutina(nombre string, nuevosEjerciciosTotales []*Ejercicio) error {
	// Validar que la rutina exista
	if _, existe := lista.listaDeRutinas[nombre]; !existe {
		return errors.New("la rutina no existe")
	}
	duracionRutina := calcularDuracion(nuevosEjerciciosTotales)
	caloriasRutina := calcularCaloriasTotales(nuevosEjerciciosTotales)
	tipoEjerciciosRutina := calcularTipoEjercicios(nuevosEjerciciosTotales)
	dificultadRutina := calcularDificultadEjercicios(nuevosEjerciciosTotales)
	rutina := &Rutina{
		nombre:                  nombre,
		duracion:                duracionRutina,
		ejerciciosTotales:       nuevosEjerciciosTotales,
		tipoDeEjercicios:        tipoEjerciciosRutina,
		caloriasQuemadasTotales: caloriasRutina,
		dificultad:              dificultadRutina,
	}
	lista.listaDeRutinas[nombre] = rutina
	return nil
}

// ListarRutinas permite listar todas las rutinas contenidas dentro del map
// de listaDeRutinas
func (lista *ListaDeRutinas) ListarRutinas() ([]*Rutina, error) {
	if len(lista.listaDeRutinas) == 0 {
		return nil, errors.New("no hay ninguna rutina para listar")
	}
	rutinas := make([]*Rutina, 0, len(lista.listaDeRutinas))
	for _, rutina := range lista.listaDeRutinas {
		rutinas = append(rutinas, rutina)
	}
	return rutinas, nil
}

// Método para agregar un ejercicio al map de ejerciciosTotales de una Rutina en particular
func (lista *ListaDeRutinas) AgregarEjercicioARutina(nombre string, ejercicio *Ejercicio) error {
	// Verificar si la rutina existe en la lista
	rutina, error := lista.ConsultarRutina(nombre)
	if error != nil {
		return error
	}
	// Verificar si el ejercicio ya está dentro del slice de la Rutina
	for _, ejer := range rutina.ejerciciosTotales {
		if ejer.nombre == ejercicio.nombre {
			return errors.New("el ejercicio ya está dentro de la rutina")
		}
	}
	// Agregar el ejercicio al slice de ejerciciosTotales de la Rutina
	rutina.ejerciciosTotales = append(rutina.ejerciciosTotales, ejercicio)
	// Actualizar la rutina en el map de ListaDeRutinas
	lista.listaDeRutinas[nombre] = rutina
	return nil
}

// Método para eliminar un ejercicio al map de ejerciciosTotales de una Rutina en particular
func (lista *ListaDeRutinas) EliminaEjercicioDeRutina(nombre string, ejercicio *Ejercicio) error {
	// Verificar si la rutina existe en la lista
	rutina, error := lista.ConsultarRutina(nombre)
	if error != nil {
		return error
	}
	// Encontrar el índice dentro del slice donde está el Ejercicio buscado
	indice := -1
	for i := range rutina.ejerciciosTotales {
		if rutina.ejerciciosTotales[i].nombre == ejercicio.nombre {
			indice = i
		}
	}
	// Si no encontró el ejercicio dentro del slice, devolver un error
	if indice == -1 {
		return errors.New("el ejercicio no existe dentro de la rutina")
	}
	// Si encontró el ejercicio, se crea un nuevo slice con todos los elementos menos
	// el identificado
	nuevosEjerciciosTotales := make([]*Ejercicio, 0)

	for _, ejer := range rutina.ejerciciosTotales {
		if ejer.nombre != rutina.ejerciciosTotales[indice].nombre {
			nuevosEjerciciosTotales = append(nuevosEjerciciosTotales, ejer)
		}
	}
	// Se reemplaza el slice existente por el nuevo
	rutina.ejerciciosTotales = nuevosEjerciciosTotales

	return nil
}

// Heapsort para ordenamiento por duración
func HeapSort(ejercicios []*Ejercicio) []*Ejercicio {
	size := len(ejercicios)
	heapify(ejercicios) // Construye un heap máximo a partir del arreglo

	// En cada iteración, se extrae el elemento máximo del heap y se coloca al final del arreglo.
	// Luego, se ajusta el heap hacia abajo para mantener la propiedad del heap.
	end := size - 1
	for end > 0 {
		// Intercambia el máximo actual con el último elemento del arreglo
		ejercicios[end], ejercicios[0] = ejercicios[0], ejercicios[end]
		// Ajusta el heap hacia abajo (restaura la propiedad del heap)
		downHeap(ejercicios, 0, end-1)
		// Reduce el tamaño efectivo del arreglo en 1 para excluir el elemento ya ordenado
		end--
	}
	return ejercicios
}

func heapify(ejercicios []*Ejercicio) {
	size := len(ejercicios)
	// El primer nodo que tiene hijos se encuentra en la posición (size - 2) / 2.
	start := (size - 2) / 2

	// Comienza desde el último padre y ajusta cada subárbol hacia abajo para cumplir la propiedad del heap.
	for start >= 0 {
		downHeap(ejercicios, start, size-1)
		start--
	}
}

func downHeap(ejercicios []*Ejercicio, start, end int) {
	father := start
	leftSon := father*2 + 1
	rightSon := leftSon + 1

	// Mientras el padre tenga al menos un hijo
	for leftSon <= end {
		// Si el padre tiene dos hijos, nos quedamos con el menor
		if rightSon <= end && ejercicios[rightSon].tiempo < ejercicios[leftSon].tiempo {
			leftSon = rightSon
		}
		// Si el hijo es menor que el padre, los intercambiamos
		if ejercicios[leftSon].tiempo < ejercicios[father].tiempo {
			ejercicios[leftSon], ejercicios[father] = ejercicios[father], ejercicios[leftSon]
			// El hijo se convierte en el padre
			father = leftSon
			leftSon = father*2 + 1
			rightSon = leftSon + 1
		} else {
			return
		}
	}
}

// Generación Automágica de Rutinas 1
func (lista *ListaDeRutinas) GeneracionAutomagica(nombre string, duracionTotal int, tipo string, dificultad string, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Filtrar los ejercicios que cumplan con los criterios especificados
	ejerciciosFiltrados, err := listaEjercicios.FiltrarEjercicios(tipo, dificultad, 0)
	if err != nil {
		return nil, err
	}

	// Ordenar los ejercicios filtrados por duración ascendente usando HeapSort
	ejerciciosOrdenados := HeapSort(ejerciciosFiltrados)

	var rutinaEjerciciosOrdenados []*Ejercicio
	tiempoAcumulado := 0
	i := 0

	// Seleccionar los ejercicios de manera greedy
	for i < len(ejerciciosOrdenados) && tiempoAcumulado < duracionTotal {
		ejercicio := ejerciciosOrdenados[i]
		if tiempoAcumulado+ejercicio.tiempo <= duracionTotal {
			rutinaEjerciciosOrdenados = append(rutinaEjerciciosOrdenados, ejercicio)
			tiempoAcumulado += ejercicio.tiempo
		}
		i++
	}

	if len(rutinaEjerciciosOrdenados) == 0 {
		return nil, errors.New("no se pudieron seleccionar ejercicios para la duración total especificada")
	}
	if tiempoAcumulado < duracionTotal {
		return nil, errors.New("no se puede alcanzar el tiempo deseado con los ejercicios existentes")
	}

	// Crear y agregar la rutina a la lista de rutinas
	// Usar AgregarRutina
	rutina := &Rutina{
		nombre:                  nombre,
		duracion:                tiempoAcumulado,
		ejerciciosTotales:       rutinaEjerciciosOrdenados,
		tipoDeEjercicios:        calcularTipoEjercicios(rutinaEjerciciosOrdenados),
		caloriasQuemadasTotales: calcularCaloriasTotales(rutinaEjerciciosOrdenados),
		dificultad:              calcularDificultadEjercicios(rutinaEjerciciosOrdenados),
	}

	lista.listaDeRutinas[nombre] = rutina
	return rutina, nil
}

// Generación Automágica de Rutinas 2

func (lista *ListaDeRutinas) GeneracionAutomagica2(nombre string, caloriasObjetivo int, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Obtener todos los ejercicios
	ejerciciosDisponibles, err := listaEjercicios.ListarEjercicios()
	if err != nil {
		return nil, err
	}

	// Ordenar los ejercicios por duración usando HeapSort
	ejerciciosOrdenados := HeapSort(ejerciciosDisponibles)

	var rutinaEjercicios []*Ejercicio
	caloriasAcumuladas := 0

	for _, ejercicio := range ejerciciosOrdenados {
		// Verificar si agregar el ejercicio excede las calorías objetivo
		if caloriasAcumuladas+ejercicio.calorias <= caloriasObjetivo {
			// Agregar el ejercicio a la rutina
			rutinaEjercicios = append(rutinaEjercicios, ejercicio)
			caloriasAcumuladas += ejercicio.calorias

			// Verificar si se alcanzaron las calorías objetivo
			if caloriasAcumuladas == caloriasObjetivo {
				break
			}
		}
	}

	// Verificar si se pudieron seleccionar ejercicios para alcanzar las calorías objetivo
	if len(rutinaEjercicios) == 0 {
		return nil, errors.New("no se pudieron seleccionar ejercicios para alcanzar las calorías objetivo")
	}
	if caloriasAcumuladas < caloriasObjetivo {
		return nil, errors.New("no se puede alcanzar las calorías deseadas con los ejercicios existentes")
	}

	// Crear y agregar la rutina a la lista de rutinas
	// Usar AgregarRutina
	rutina := &Rutina{
		nombre:                  nombre,
		duracion:                calcularDuracion(rutinaEjercicios),
		ejerciciosTotales:       rutinaEjercicios,
		tipoDeEjercicios:        calcularTipoEjercicios(rutinaEjercicios),
		caloriasQuemadasTotales: caloriasAcumuladas,
		dificultad:              calcularDificultadEjercicios(rutinaEjercicios),
	}

	lista.listaDeRutinas[nombre] = rutina
	return rutina, nil
}

// Se deberá implementar un algoritmo que permita generar automáticamente una rutina nueva, teniendo en cuenta los siguientes parámetros:

// Nombre de la rutina
// Tipo de puntos a maximizar (cardio, fuerza, flexibilidad)
// Duración máxima de la rutina
// El algoritmo deberá seleccionar los ejercicios que cumplan con los parámetros establecidos y que maximicen el puntaje de la rutina en la dimensión solicitada. Para ello, se empleará el concepto de "duración fija, máximos puntos". No pueden repetirse ejercicios en la rutina.
// Para ello, se empleará el concepto de "duración fija, máximos puntos". No pueden repetirse ejercicios en la rutina.
func (lista *ListaDeRutinas) GeneracionAutomagica3(nombre string, duracionFija int, tipoDeEjDeMaximosPuntos string, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Filtrar los ejercicios que cumplan con los criterios especificados
	ejerciciosFiltrados, err := listaEjercicios.FiltrarEjercicios(tipoDeEjDeMaximosPuntos, "", 0)
	if err != nil {
		return nil, err
	}

	var rutinaEjerciciosPorTipo []*Ejercicio ///esta sera la lista temporal para ir agregando los ejercicios
	tiempo := 0
	i := 0

	for i < len(ejerciciosFiltrados) && ejerciciosFiltrados[i].tiempo < duracionFija {

		ejercicio := ejerciciosFiltrados[i]
		//SE CONTROLA QUE NO SUPERE AL TPO PASADO COMO PARAMETRO// SI NO LO SUPERA, LO AGREGA A LA RUTINA
		if tiempo+ejercicio.tiempo <= duracionFija {
			rutinaEjerciciosPorTipo = append(rutinaEjerciciosPorTipo, ejercicio)
			tiempo += ejercicio.tiempo
		}
		i++
	}
	if tiempo < duracionFija {
		return nil, errors.New("no se puede alcanzar el tiempo deseado con los ejercicios existentes")
	}

	Backtracking(duracionFija, ejerciciosFiltrados, rutinaEjerciciosPorTipo)

	// Crear y agregar la rutina a la lista de rutinas
	rutina := &Rutina{
		nombre:                  nombre,
		duracion:                tiempo,
		ejerciciosTotales:       rutinaEjerciciosPorTipo,
		tipoDeEjercicios:        calcularTipoEjercicios(rutinaEjerciciosPorTipo),
		caloriasQuemadasTotales: calcularCaloriasTotales(rutinaEjerciciosPorTipo),
		dificultad:              calcularDificultadEjercicios(rutinaEjerciciosPorTipo),
	}

	lista.listaDeRutinas[nombre] = rutina
	return rutina, err
}

func Backtracking(duracionFija int, ejerciciosFiltrados []*Ejercicio, rutinaEjerciciosPorTipo []*Ejercicio) {

	if len(ejerciciosFiltrados) == 1 && ejerciciosFiltrados[0].tiempo > duracionFija {
		println(errors.New("no se puede alcanzar el tiempo deseado con los ejercicios existentes"))
	}

	if len(ejerciciosFiltrados) > 0 && ejerciciosFiltrados[0].tiempo > duracionFija {
		println(errors.New("no se puede alcanzar el tiempo deseado con los ejercicios existentes"))
	}

	tiempo := 0
	if len(ejerciciosFiltrados) > 0 && ejerciciosFiltrados[0].tiempo <= duracionFija {

		rutinaEjerciciosPorTipo = append(rutinaEjerciciosPorTipo, ejerciciosFiltrados[0])
		tiempo += ejerciciosFiltrados[0].tiempo
		ejerciciosFiltrados = append(ejerciciosFiltrados[:0], ejerciciosFiltrados[1:]...)

		Backtracking(duracionFija, ejerciciosFiltrados, rutinaEjerciciosPorTipo)

	}

}

// Versión 2 de Automagicas3
func (lista *ListaDeRutinas) GeneracionAutomagica3v2(nombre string, duracionTotal int, tipo string, listaEjercicios *ListaDeEjercicios) (*Rutina, error) {
	// Filtrar los ejercicios por tipo
	ejerciciosFiltrados, err := listaEjercicios.FiltrarEjercicios(tipo, "", 0)
	if err != nil {
		return nil, err
	}

	// Crear una tabla para almacenar los máximos puntos
	n := len(ejerciciosFiltrados)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, duracionTotal+1)
	}

	// Llenar la tabla de forma dinámica
	for i := 1; i <= n; i++ {
		ejercicio := ejerciciosFiltrados[i-1]
		puntos := 0
		for j, t := range ejercicio.tipoDeEjercicio {
			if t == tipo {
				puntos = ejercicio.puntosPorTipoDeEjercicio[j]
				break
			}
		}
		for j := 1; j <= duracionTotal; j++ {
			if ejercicio.tiempo <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-ejercicio.tiempo]+puntos)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// Recuperar los ejercicios seleccionados
	tiempoRestante := duracionTotal
	rutinaEjercicios := []*Ejercicio{}
	for i := n; i > 0 && tiempoRestante > 0; i-- {
		if dp[i][tiempoRestante] != dp[i-1][tiempoRestante] {
			ejercicio := ejerciciosFiltrados[i-1]
			rutinaEjercicios = append(rutinaEjercicios, ejercicio)
			tiempoRestante -= ejercicio.tiempo
		}
	}

	// Validar que se pudieron seleccionar ejercicios
	if len(rutinaEjercicios) == 0 {
		return nil, errors.New("no se pudieron seleccionar ejercicios")
	}

	// Crear y agregar la rutina
	rutina := &Rutina{
		nombre:                  nombre,
		duracion:                calcularDuracion(rutinaEjercicios),
		ejerciciosTotales:       rutinaEjercicios,
		tipoDeEjercicios:        tipo,
		caloriasQuemadasTotales: calcularCaloriasTotales(rutinaEjercicios),
		dificultad:              calcularDificultadEjercicios(rutinaEjercicios),
	}

	lista.listaDeRutinas[nombre] = rutina
	return rutina, nil
}
