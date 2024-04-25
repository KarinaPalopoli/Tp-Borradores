package rutinaDeEjercicios

import (
	"errors"
)

// Estructura de Rutina
type Rutina struct {
	nombre string
	duracion int
	ejerciciosTotales []*Ejercicio
	tipoDeEjercicios string // fuerza, cardio, balance
	caloriasQuemadasTotales int
	dificultad string // principiante, intermedio, avanzado
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
		for _, dificultad := range ejercicio.dificultad {
			dificultades[string(dificultad)]++
		}
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
	return &ListaDeRutinas{listaDeRutinas: make(map[string]*Rutina)}
}

// AgregarRutina crea una rutina y la agrega al map de listaDeRutinas
func (lista *ListaDeRutinas) AgregarRutina(nombre string,	ejerciciosTotales []*Ejercicio) error {
    // Validar que la rutina no exista previamente
    if _, existe := lista.listaDeRutinas[nombre]; existe {
        return errors.New("la rutina ya existe")
    }
	if len(ejerciciosTotales) == 0{
		return errors.New("una rutina debe contener al menos 1 ejercicio")
	}
	duracionRutina:= calcularDuracion(ejerciciosTotales)
	caloriasRutina:= calcularCaloriasTotales(ejerciciosTotales)
	tipoEjerciciosRutina:= calcularTipoEjercicios(ejerciciosTotales)
	dificultadRutina:= calcularDificultadEjercicios(ejerciciosTotales)
	rutina := &Rutina{
        nombre: nombre,
        duracion: duracionRutina,
        ejerciciosTotales: ejerciciosTotales,
        tipoDeEjercicios: tipoEjerciciosRutina,
        caloriasQuemadasTotales: caloriasRutina,
        dificultad: dificultadRutina,
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
		rutina, existe := lista.listaDeRutinas[nombre];
		if !existe{
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
		duracionRutina:= calcularDuracion(nuevosEjerciciosTotales)
		caloriasRutina:= calcularCaloriasTotales(nuevosEjerciciosTotales)
		tipoEjerciciosRutina:= calcularTipoEjercicios(nuevosEjerciciosTotales)
		dificultadRutina:= calcularDificultadEjercicios(nuevosEjerciciosTotales)
		rutina := &Rutina{
        nombre: nombre,
		duracion: duracionRutina,
        ejerciciosTotales: nuevosEjerciciosTotales,
        tipoDeEjercicios: tipoEjerciciosRutina,
        caloriasQuemadasTotales: caloriasRutina,
        dificultad: dificultadRutina,
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
	if error != nil{
		return error
	}
	// Verificar si el ejercicio ya está dentro del slice de la Rutina
	for _,ejer:= range rutina.ejerciciosTotales{
		if ejer.nombre == ejercicio.nombre{
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
	if error != nil{
		return error
	}
	// Encontrar el índice dentro del slice donde está el Ejercicio buscado
	indice:= -1
	for i:= range rutina.ejerciciosTotales{
		if rutina.ejerciciosTotales[i].nombre == ejercicio.nombre{
			indice = i
		}
	}
	// Si no encontró el ejercicio dentro del slice, devolver un error
	if indice == -1{
		return errors.New("el ejercicio no existe dentro de la rutina")
	}
	// Si encontró el ejercicio, se crea un nuevo slice con todos los elementos menos
	// el identificado
	nuevosEjerciciosTotales:= make([]*Ejercicio,0)

	for _,ejer:= range rutina.ejerciciosTotales{
		if ejer.nombre != rutina.ejerciciosTotales[indice].nombre{
			nuevosEjerciciosTotales = append(nuevosEjerciciosTotales, ejer)
		}
	}
	// Se reemplaza el slice existente por el nuevo
	rutina.ejerciciosTotales = nuevosEjerciciosTotales

	return nil
}