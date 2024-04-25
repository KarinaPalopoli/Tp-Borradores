package rutinaDeEjercicios

import (
	"errors"
)

// Estructura de Ejercicio
type Ejercicio struct {
	nombre string
	descripcion string
	tiempo int
	calorias int
	tipoDeEjercicio []string // fuerza, cardio, balance
	puntosPorTipoDeEjercicio []int
	dificultad string // principiante, intermedio, avanzado
}

// Estructura para almacenar los ejercicios
type ListaDeEjercicios struct {
	listaDeEjercicios map[string]*Ejercicio
}

// Inicializa una ListaDeEjercicios y crea el map vacío
func NewListaDeEjercicios() *ListaDeEjercicios {
	return &ListaDeEjercicios{listaDeEjercicios: make(map[string]*Ejercicio)}
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
        nombre: nombre,
        descripcion: descripcion,
        tiempo: tiempo,
        calorias: calorias,
        tipoDeEjercicio: tipoDeEjercicio,
        puntosPorTipoDeEjercicio: puntosPorTipoDeEjercicio,
        dificultad: dificultad,
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
    if !existe{
        return nil, errors.New("el ejercicio no existe")
    } 
    return ejercicio, nil
}

// FiltrarEjercicios permite filtrar los ejercicios que cumplan con los criterios indicados por parámetro
// y devuleve un slice con los ejercicios que cumplan
func (lista *ListaDeEjercicios) FiltrarEjercicios(tipo string, dificultad string, minCalorias int) ([]*Ejercicio, error) {
    ejerciciosFiltrados:= make([]*Ejercicio, 0)
    // Recorrer todos los ejercicios
    for _,ejercicio := range lista.listaDeEjercicios{
        // Se crea un booleano para ver si el ejercicio cumple los filtros o no
        // se inicializa en true y luego las comprobaciones van pasando a false los que no cumplan
        cumpleFiltro:= true
        // Se verifica el tipo de ejercicio, si es que se pasa por parámetro
        if tipo != ""{
            tipoEncontrado:= false
            for _,t:= range ejercicio.tipoDeEjercicio{
                if t == tipo{
                    tipoEncontrado = true
                }
            }
            if !tipoEncontrado {
                cumpleFiltro = false                
            }
        }
        // Se verifica la dificultad, si es que se pasa por parámetro
        if dificultad != "" && ejercicio.dificultad != dificultad{
            cumpleFiltro = false
        }
        // Se verifican las calorías mínimas, si es que se pasa por parámetro
        if minCalorias > 0 && ejercicio.calorias < minCalorias{
            cumpleFiltro = false
        }
        // Si el ejercicio pasa los filtros, se agrega al slice
        if cumpleFiltro{
            ejerciciosFiltrados = append(ejerciciosFiltrados, ejercicio)
        }
    }
    // Chequeamos que el slice tenga algún elemento o esté vacío
    if len(ejerciciosFiltrados)== 0{
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
    if !existe{
        return errors.New("el ejercicio no existe")
    } 

    ejercicio := &Ejercicio{
        nombre: nombre,
        descripcion: nuevaDescripcion,
        tiempo: nuevoTiempo,
        calorias: nuevasCalorias,
        tipoDeEjercicio: nuevoTipoDeEjercicio,
        puntosPorTipoDeEjercicio: nuevosPuntosPorTipoDeEjercicio,
        dificultad: nuevaDificultad,
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