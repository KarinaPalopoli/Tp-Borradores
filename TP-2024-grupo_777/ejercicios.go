package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

// Estructura de Ejercicio
// Debe considerarse que los ejercicios tendrán etiquetas para definir su tipo de ejercicio y dificultad. Las etiquetas serán un conjunto de palabras clave que permitirán clasificar los ejercicios. Por ejemplo, un ejercicio de sentadillas podría tener las etiquetas "fuerza" y "piernas". Sin embargo, las dificultades son únicas; por ejemplo, las sentadillas serán de dificultad "media" sólamente.
type Ejercicio struct {
	nombre                   string
	descripcion              string
	tiempo                   int
	calorias                 int
	tipoDeEjercicio          []string // fuerza, cardio, balance
	puntosPorTipoDeEjercicio []int
	dificultad               Dificultad // principiante, intermedio, avanzado
}

// Definir una estructura para representar los datos del CSV
// no se puede usar el struc Ejercicio porque para que gocsv guarde los datos deben ser public.
type EjercicioCSV struct {
	Nombre                   string `csv:"nombre"`
	Descripcion              string `csv:"descripcion"`
	Tiempo                   int    `csv:"tiempo"`
	Calorias                 int    `csv:"calorias"`
	TipoDeEjercicio          string `csv:"tipoDeEjercicio"`
	PuntosPorTipoDeEjercicio string `csv:"puntosPorTipoDeEjercicio"`
	Dificultad               string `csv:"dificultad"`
}

type Dificultad string

const (
	Principiante Dificultad = "principiante"
	Intermedio   Dificultad = "intermedio"
	Avanzado     Dificultad = "avanzado"
)

// Estructura para almacenar los ejercicios
type ListaDeEjercicios struct {
	listaDeEjercicios map[string]*Ejercicio
}

// Inicializa una ListaDeEjercicios y crea el map vacío
func NewListaDeEjercicios() *ListaDeEjercicios {
	listaTemp := &ListaDeEjercicios{
		listaDeEjercicios: make(map[string]*Ejercicio)}

		
		
			// Iterar sobre los ejercicios para actualizar el mayor puntaje por tipo de ejercicio
			for _, ejercicio := range listaTemp.listaDeEjercicios {
				// Crear un mapa temporal para almacenar el puntaje máximo por tipo de ejercicio
				maxPuntosPorTipoDeEjercicio := make(map[string]int)
		
				// Actualizar el puntaje máximo por tipo de ejercicio
				for i, tipo := range ejercicio.tipoDeEjercicio {
					if puntos := ejercicio.puntosPorTipoDeEjercicio[i]; puntos > maxPuntosPorTipoDeEjercicio[tipo] {
						maxPuntosPorTipoDeEjercicio[tipo] = puntos
					}
				}
		
				// Actualizar los puntajes por tipo de ejercicio en el ejercicio
				for i, tipo := range ejercicio.tipoDeEjercicio {
					ejercicio.puntosPorTipoDeEjercicio[i] = maxPuntosPorTipoDeEjercicio[tipo]
				}
			}
		
	err := listaTemp.cargarEjerciciosDesdeCSV()
	if err != nil {
		log.Fatal("Error: ", err)

	}

	return listaTemp
}

// Función para cargar ejercicios desde un archivo CSV usando gocsv
func (lista *ListaDeEjercicios) cargarEjerciciosDesdeCSV() error {
	// Abrir el archivo CSV
	file, err := os.OpenFile("ejercicios.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// Crear un slice para almacenar los ejercicios
	var ejercicios []*EjercicioCSV

	// Leer el archivo CSV usando gocsv
	if err := gocsv.UnmarshalFile(file, &ejercicios); err != nil {
		return err
	}

	// Convertir los datos cargados a la estructura de Ejercicio y agregarlos a una lista de ejercicios
	for _, ecsv := range ejercicios {
		dificultad := Dificultad(ecsv.Dificultad)
		// Crear un nuevo ejercicio
		ejercicio := &Ejercicio{
			nombre:                   ecsv.Nombre,
			descripcion:              ecsv.Descripcion,
			tiempo:                   ecsv.Tiempo,
			calorias:                 ecsv.Calorias,
			tipoDeEjercicio:          strings.Split(ecsv.TipoDeEjercicio, ","),
			puntosPorTipoDeEjercicio: lista.stringToIntArray(ecsv.PuntosPorTipoDeEjercicio),
			dificultad:               dificultad,
		}

		lista.listaDeEjercicios[ecsv.Nombre] = ejercicio
	}

	return nil
}

func (lista *ListaDeEjercicios) stringToIntArray(str string) []int {
	intSlice := make([]int, 0)
	for _, s := range strings.Split(str, ",") {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Error convirtiendo a entero: %v\n", err)
			return nil
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}

// Función para convertir un ejercicio a EjercicioCSV
func ejercicioToEjercicioCSV(e *Ejercicio) *EjercicioCSV {

	strPuntosPorTipoDeEjercicio := make([]string, len(e.puntosPorTipoDeEjercicio))
	for i, num := range e.puntosPorTipoDeEjercicio {
		strPuntosPorTipoDeEjercicio[i] = strconv.Itoa(num)
	}

	return &EjercicioCSV{
		Nombre:                   e.nombre,
		Descripcion:              e.descripcion,
		Tiempo:                   e.tiempo,
		Calorias:                 e.calorias,
		TipoDeEjercicio:          strings.Join(e.tipoDeEjercicio, ","),
		PuntosPorTipoDeEjercicio: strings.Join(strPuntosPorTipoDeEjercicio, ","),
		Dificultad:               string(e.dificultad),
	}
}

// Función para guardar ejercicios en un archivo CSV usando gocsv
func (lista *ListaDeEjercicios) guardarEjerciciosEnCSV() error {
	// Abrir el archivo CSV para escritura
	file, err := os.OpenFile("ejercicios.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convertir los ejercicios a la estructura EjercicioCSV
	var ejerciciosCSV []*EjercicioCSV
	for _, ejercicio := range lista.listaDeEjercicios {
		ejerciciosCSV = append(ejerciciosCSV, ejercicioToEjercicioCSV(ejercicio))
	}

	// Escribir al archivo CSV usando gocsv
	if err := gocsv.MarshalFile(&ejerciciosCSV, file); err != nil {
		return err
	}

	fmt.Println("\n\n--------------------------------")
	fmt.Printf("Base de Datos CSV guardada existosamente\n\n")

	return nil
}

// AgregarEjercicio crea un ejercicio y lo agrega al map de listaDeEjercicios
func (lista *ListaDeEjercicios) AgregarEjercicio(nombre string, descripcion string, tiempo int, calorias int, tipoDeEjercicio []string, puntosPorTipoDeEjercicio []int, dificultad string) error {
	// Validar que el ejercicio no exista previamente
	if _, existe := lista.listaDeEjercicios[nombre]; existe {
		return errors.New("el ejercicio ya existe")
	}
	// Validar la longitud de los slices tipoDeEjercicio y puntosPorTipoDeEjercicio
	if len(tipoDeEjercicio) != len(puntosPorTipoDeEjercicio) {
		return errors.New("los slices de tipoDeEjercicio y puntosPorTipoDeEjercicio deben tener la misma longitud")
	}

	ejercicio := &Ejercicio{
		nombre:                   nombre,
		descripcion:              descripcion,
		tiempo:                   tiempo,
		calorias:                 calorias,
		tipoDeEjercicio:          tipoDeEjercicio,
		puntosPorTipoDeEjercicio: puntosPorTipoDeEjercicio,
		dificultad:               Dificultad(dificultad),
	}
	lista.listaDeEjercicios[nombre] = ejercicio
	return nil
}

// BorrarEjercicio elimina el par key value, a partir de la key indicada
func (lista *ListaDeEjercicios) BorrarEjercicio(nombre string) error {
	// Validar que el ejercicio exista
	if _, existe := lista.listaDeEjercicios[nombre]; !existe {
		return errors.New("el ejercicio no existe")
	}
	delete(lista.listaDeEjercicios, nombre)
	return nil
}

// ConsultarEjercicioPorNombre busca el ejercicio a partir del nombre indicado y devuelve el Ejercicio
func (lista *ListaDeEjercicios) ConsultarEjercicioPorNombre(nombre string) (*Ejercicio, error) {
	// Validar que el ejercicio exista
	ejercicio, existe := lista.listaDeEjercicios[nombre]
	if !existe {
		return ejercicio, errors.New("el ejercicio no existe")
	}
	return ejercicio, nil
}

// FiltrarEjercicios permite filtrar los ejercicios que cumplan con los criterios indicados por parámetro
// y devuleve un slice con los ejercicios que cumplan
func (lista *ListaDeEjercicios) FiltrarEjercicios(tipo string, dificultad string, minCalorias int) ([]*Ejercicio, error) {
	ejerciciosFiltrados := make([]*Ejercicio, 0)
	// Recorrer todos los ejercicios
	for _, ejercicio := range lista.listaDeEjercicios {
		// Se crea un booleano para ver si el ejercicio cumple los filtros o no
		// se inicializa en true y luego las comprobaciones van pasando a false los que no cumplan
		cumpleFiltro := true
		// Se verifica el tipo de ejercicio, si es que se pasa por parámetro
		if tipo != "" {
			tipoEncontrado := false
			for _, t := range ejercicio.tipoDeEjercicio {
				if t == tipo {
					tipoEncontrado = true
				}
			}
			if !tipoEncontrado {
				cumpleFiltro = false
			}
		}
		// Se verifica la dificultad, si es que se pasa por parámetro
		if dificultad != "" && ejercicio.dificultad != Dificultad(dificultad) {
			cumpleFiltro = false
		}
		// Se verifican las calorías mínimas, si es que se pasa por parámetro
		if minCalorias > 0 && ejercicio.calorias < minCalorias {
			cumpleFiltro = false
		}
		// Si el ejercicio pasa los filtros, se agrega al slice
		if cumpleFiltro {
			ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
		}
	}
	// Chequeamos que el slice tenga algún elemento o esté vacío
	if len(ejerciciosFiltrados) == 0 {
		return nil, errors.New("no hay ejercicios que cumplan esas condiciones")
	}
	return ejerciciosFiltrados, nil
}

// ModificarEjercicio permite modificar los valores de un ejercicio,
// a partir de identificar al mismo a partir de la key indicada
func (lista *ListaDeEjercicios) ModificarEjercicio(nombre string, nuevaDescripcion string, nuevoTiempo int, nuevasCalorias int, nuevoTipoDeEjercicio []string, nuevosPuntosPorTipoDeEjercicio []int, nuevaDificultad string) error {
	// Validar la longitud de los slices tipoDeEjercicio y puntosPorTipoDeEjercicio
	if len(nuevoTipoDeEjercicio) != len(nuevosPuntosPorTipoDeEjercicio) {
		return errors.New("los slices de tipoDeEjercicio y puntosPorTipoDeEjercicio deben tener la misma longitud")
	}
	// Validar que el ejercicio exista
	_, existe := lista.listaDeEjercicios[nombre]
	if !existe {
		return errors.New("el ejercicio no existe")
	}

	ejercicio := &Ejercicio{
		nombre:                   nombre,
		descripcion:              nuevaDescripcion,
		tiempo:                   nuevoTiempo,
		calorias:                 nuevasCalorias,
		tipoDeEjercicio:          nuevoTipoDeEjercicio,
		puntosPorTipoDeEjercicio: nuevosPuntosPorTipoDeEjercicio,
		dificultad:               Dificultad(nuevaDificultad),
	}
	lista.listaDeEjercicios[nombre] = ejercicio
	return nil
}

// ListarEjercicios permite listar todos los ejercicios contenidos dentro del map
// de listaDeEjercicios
func (lista *ListaDeEjercicios) ListarEjercicios() ([]*Ejercicio, error) {
	if len(lista.listaDeEjercicios) == 0 {
		return nil, errors.New("no hay ningún ejercicio para listar")
	}
	ejercicios := make([]*Ejercicio, 0, len(lista.listaDeEjercicios))
	for _, ejercicio := range lista.listaDeEjercicios {
		ejercicios = append(ejercicios, ejercicio)
	}
	return ejercicios, nil
}
