package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularDuracion(t *testing.T) {
	// Prueba con Lista de ejercicios vacía
	ejercicios := []*Ejercicio{}
	duracion := calcularDuracion(ejercicios)
	assert.Equal(t, 0, duracion, "La duración debe ser 0 para una lista vacía")

	// Prueba con Lista de ejercicios con elementos
	ejercicios = []*Ejercicio{
		{
			nombre: "Flexiones",
			descripcion: "",
			tiempo:      30,
			calorias:                 0,
			tipoDeEjercicio:          []string{},
			puntosPorTipoDeEjercicio: []int{},
			dificultad:               "",
		},
		{
			nombre: "Sentadillas",
			descripcion: "",
			tiempo:      45,
			calorias:                 0,
			tipoDeEjercicio:          []string{},
			puntosPorTipoDeEjercicio: []int{},
			dificultad:               "",
		},
		{
			nombre: "Abdominales",
			descripcion: "",
			tiempo:      20,
			calorias:                 0,
			tipoDeEjercicio:          []string{},
			puntosPorTipoDeEjercicio: []int{},
			dificultad:               "",
		},
	}
	duracionEsperada := 95
	duracion = calcularDuracion(ejercicios)
	assert.Equal(t, duracionEsperada, duracion, "La duración coincide con la esperada")
}

// COMPRUEBA LA FUNCION QUE CALCULA EL TIPO DE EJERCICIO DE UNA RUTINA
// TENIENDO EN CUENTA QUE EL MAS FRECUENTE ES “FUERZA”
func TestCalcularTipoEjerciciosConMultiplicidadEnFuerza(t *testing.T) {
	ejercicios := []*Ejercicio{
		{tipoDeEjercicio: []string{"Cardio"}},
		{tipoDeEjercicio: []string{"Fuerza"}},
		{tipoDeEjercicio: []string{"Fuerza"}},
		{tipoDeEjercicio: []string{"Balance"}},
	}

	resultado := calcularTipoEjercicios(ejercicios)

	assert.Equal(t, "Fuerza", resultado)
}

// VERIFICA EL TEST DE CALCULAR A QUÉ TIPO DE EJERRCICIO SE
// LE ASIGNA A UNA RUTINA CUANDO CARDIO y FUERZA SE REPITEN LA MISMA CANTIDAD DE VECES
func TestCalculaTipoEjercicioConCantidadtotalIgualEnDosTiposRepetidos(t *testing.T) {
	ejercicios := []*Ejercicio{
		{tipoDeEjercicio: []string{"Cardio"}},
		{tipoDeEjercicio: []string{"Fuerza"}},
		{tipoDeEjercicio: []string{"Balance"}},
		{tipoDeEjercicio: []string{"Cardio"}},
		{tipoDeEjercicio: []string{"Fuerza"}},
	}

	resultado := calcularTipoEjercicios(ejercicios)
	if resultado != "Cardio" && resultado != "Fuerza" {
		t.Errorf("Esperado 'Cardio' or Fuerza, se obtuvo '%s'", resultado)
	}
}

func TestAgregarRutina(t *testing.T) {
	lista := NewListaDeRutinas()

	// prueba con una rutina válida
	ejerciciosA := []*Ejercicio{
		{
			nombre: "Flexiones",
			descripcion: "",
			tiempo:      30,
			calorias: 100,
			tipoDeEjercicio: []string{"fuerza"},
			puntosPorTipoDeEjercicio: []int{},
			dificultad:  "principiante",
		},
	
	}
	err := lista.AgregarRutina("Rutina1",ejerciciosA)
	assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

	// prueba con nombre de rutina duplicado
	err = lista.AgregarRutina("Rutina1", ejerciciosA)
	assert.Error(t, err, "Se esperaba un error al intentar agregar una rutina con nombre duplicado")
	assert.EqualError(t, err, "la rutina ya existe")

	// Agregando rutina sin ejercicios
	err = lista.AgregarRutina("Rutina2", []*Ejercicio{})
	assert.Error(t, err, "Se esperaba un error al intentar agregar una rutina sin ejercicios")
	assert.EqualError(t, err, "una rutina debe contener al menos 1 ejercicio")
}

func TestBorrarRutina(t *testing.T) {
	lista := NewListaDeRutinas()
	//se crea una rutina con 1 ejercicio
	ejerciciosA := []*Ejercicio{
		{nombre: "Flexiones", tiempo: 30, calorias: 100, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
		{nombre: "Sentadillas", tiempo: 45, calorias: 150, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
	}
	err := lista.AgregarRutina("Rutina1", ejerciciosA)
	assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

	// eliminando una rutina existente
	err = lista.BorrarRutina("Rutina1")
	assert.NoError(t, err, "No se esperaba un error al eliminar la rutina 1")

	// eliminando una rutina inexistente
	err = lista.BorrarRutina("Rutina1")
	assert.Error(t, err, "Se esperaba un error al intentar eliminar una rutina que no existe")
	assert.EqualError(t, err, "la rutina no existe")
}

func TestConsultarRutina(t *testing.T) {
	lista := NewListaDeRutinas()

	ejerciciosA := []*Ejercicio{
		{nombre: "Flexiones", tiempo: 30, calorias: 100, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
		{nombre: "Sentadillas", tiempo: 45, calorias: 150, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
	}
	err := lista.AgregarRutina("Rutina1", ejerciciosA)
	assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

	// consultar una rutina existente
	rutina, err := lista.ConsultarRutina("Rutina1")
	assert.NoError(t, err, "No se esperaba un error al consultar la rutina 1")
	assert.Equal(t, "Rutina1", rutina.nombre, "El nombre de la rutina no coincide")

	// intentar consultar una rutina inexistente
	_, err = lista.ConsultarRutina("Rutina2")
	assert.Error(t, err, "Se esperaba un error al intentar consultar una rutina que no existe")
	assert.EqualError(t, err, "la rutina no existe")
}

func TestModificarRutina(t *testing.T) {
	lista := NewListaDeRutinas()

	ejerciciosA := []*Ejercicio{
		{nombre: "Flexiones", tiempo: 30, calorias: 100, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
		{nombre: "Sentadillas", tiempo: 45, calorias: 150, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
	}
	err := lista.AgregarRutina("Rutina1", ejerciciosA)
	assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

	// modificando rutina existente
	nuevosEjercicios := []*Ejercicio{
		{nombre: "Abdominales", tiempo: 20},
	}
	err = lista.ModificarRutina("Rutina1", nuevosEjercicios)
	assert.NoError(t, err, "No se esperaba un error al modificar la rutina 1")

	// intentar modificar una rutina inexistente
	err = lista.ModificarRutina("Rutina2", nuevosEjercicios)
	assert.Error(t, err, "Se esperaba un error al intentar modificar una rutina que no existe")
	assert.EqualError(t, err, "la rutina no existe")
}

func TestListarRutinas(t *testing.T) {
	lista := NewListaDeRutinas()

	// listar sin rutinas
	rutinas, err := lista.ListarRutinas()
	assert.Nil(t, rutinas, "Se esperaba nil al listar rutinas cuando no hay ninguna")
	assert.Error(t, err, "Se esperaba un error al listar rutinas cuando no hay ninguna")
	assert.EqualError(t, err, "no hay ninguna rutina para listar")

	ejerciciosA := []*Ejercicio{
		{nombre: "Flexiones", tiempo: 30, calorias: 100, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
		{nombre: "Sentadillas", tiempo: 45, calorias: 150, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
	}
	err = lista.AgregarRutina("Rutina1", ejerciciosA)
	assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

	// listar con al menos una rutina
	rutinas, err = lista.ListarRutinas()
	assert.NotNil(t, rutinas, "Se esperaba una lista de rutinas al listar rutinas cuando hay al menos una")
	assert.NoError(t, err, "No se esperaba un error al listar rutinas cuando hay al menos una")
	assert.Equal(t, 1, len(rutinas), "Se esperaba una lista de rutinas con un elemento")
}

// VERIFICA EL BORRADO DE UNA RUTINA ACTUAL
func TestBorrarRutinaExistente(t *testing.T) {
	rutinas := make(map[string]*Rutina)
	lista := ListaDeRutinas{listaDeRutinas: rutinas}
	rutinas["Rutina de Lunes"] = &Rutina{nombre: "Rutina de Lunes"}

	error := lista.BorrarRutina("Rutina de Lunes")
	if error != nil {
		t.Errorf("Error al borrar una rutina existente: %s", error)
	}

	if _, existe := rutinas["Rutina de Lunes"]; existe {
		t.Error("La rutina no fue borrada correctamente")
	}
}

// VERIFICA QUE NO se puede BORRAR RUTINA INEXISTENTE
func TestBorrarRutinaNoExistente(t *testing.T) {
	rutinas := make(map[string]*Rutina)
	lista := ListaDeRutinas{listaDeRutinas: rutinas}

	error := lista.BorrarRutina("Rutina de martes")
	if error == nil || error.Error() != "la rutina no existe" {
		t.Errorf("Expected error 'la rutina no existe', got %v", error)
	}
}

// VERIFICA EN EL CASO DE AGREGAR UNA RUTINA DUPICADA
func TestAgregarEjercicioARutinaExistente(t *testing.T) {
	lista := NewListaDeRutinas()
	ejercicio := &Ejercicio{nombre: "abdominales", descripcion: "pasar de una posición tumbada a una sentada al llevar el pecho hacia los muslos", tiempo: 5, calorias: 250, tipoDeEjercicio: []string{"CARDIO"}, puntosPorTipoDeEjercicio: []int{100}, dificultad: "INTERMEDIA"}
	lista.listaDeRutinas["Rutina de Martes"] = &Rutina{
		nombre:            "Rutina de Martes",
		ejerciciosTotales: []*Ejercicio{ejercicio},
	}

	error := lista.AgregarEjercicioARutina("Rutina de Martes", ejercicio)
	if error == nil || error.Error() != "el ejercicio ya está dentro de la rutina" {
		t.Error("Expected Error")
	}
}

// VERIFICA  EJERCICIO A UNA RUTINA INEXISTENTE
func TestAgregarEjercicioARutinaInexistente(t *testing.T) {
	lista := NewListaDeRutinas()
	ejercicio := &Ejercicio{nombre: "abdominales", descripcion: "pasar de una posición tumbada a una sentada al llevar el pecho hacia los muslos", tiempo: 5, calorias: 250, tipoDeEjercicio: []string{"CARDIO"}, puntosPorTipoDeEjercicio: []int{100}, dificultad: "INTERMEDIA"}

	error := lista.AgregarEjercicioARutina("Rutina Inexistente", ejercicio)
	if error == nil || error.Error() != "la rutina no existe" {
		t.Errorf("Expected 'la rutina no existe' error, got: %v", error)
	}
}

// COMPRUEBA QUE NO SE PUEDE MODIFICAR UNA RUTINA INEXISTENTE
func TestModificarRutinaNoExistente(t *testing.T) {
	lista := NewListaDeRutinas()
	nuevoEjercicio := []*Ejercicio{
		{nombre: "abdominales", descripcion: "pasar de una posición tumbada a una sentada al llevar el pecho hacia los muslos", tiempo: 5, calorias: 250, tipoDeEjercicio: []string{"CARDIO"}, puntosPorTipoDeEjercicio: []int{100}, dificultad: "INTERMEDIA"},
	}

	error := lista.ModificarRutina("Rutina de domingos", nuevoEjercicio)
	if error == nil || error.Error() != "la rutina no existe" {
		t.Errorf("Esperado 'la rutina no existe' error, obtenido %v", error)
	}
}

// Test Generación Automágica de Rutinas
func TestGeneracionAutomagica_Exito(t *testing.T) {
	// Crear una lista de ejercicios y agregar un ejercicio
	listaEjercicios := NewListaDeEjercicios()
	err := listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripción de Flexiones de brazos", 30, 50, []string{"fuerza"}, []int{10}, "principiante")
	if err != nil {
		t.Errorf("Error al agregar ejercicio: %v", err)
	}

	// Crear una lista de rutinas
	listaDeRutinas := NewListaDeRutinas()

	// Llamar a la función GeneracionAutomagica
	rutina, err := listaDeRutinas.GeneracionAutomagica("Rutina1", 30, "fuerza", "principiante", listaEjercicios)

	// Verificar que no haya ocurrido un error
	assert.Nil(t, err, "Se esperaba que no ocurriera ningún error")

	// Verificar que la duración de la rutina sea igual a la duración total de los ejercicios disponibles
	assert.Equal(t, 30, rutina.duracion, "La duración no es igual a la duración total de los ejercicios disponibles")
}

func TestGeneracionAutomagica_Error_TipoEjercicioInexistente(t *testing.T) {
	// Crear una lista de ejercicios y agregar un ejercicio con un tipo incorrecto
	listaEjercicios := NewListaDeEjercicios()
	err := listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripción de Flexiones de brazos", 30, 50, []string{"fuerza"}, []int{10}, "principiante")
	if err != nil {
		t.Errorf("Error al agregar ejercicio: %v", err)
	}
	// Crear una lista de rutinas
	listaDeRutinas := NewListaDeRutinas()
	// Llamar a la función GeneracionAutomagica con un tipo de ejercicio inexistente
	_, err = listaDeRutinas.GeneracionAutomagica("Rutina5", 30, "cardio", "principiante", listaEjercicios)
	assert.Error(t, err, "Se esperaba un error debido al tipo de ejercicio inexistente, pero no se recibió ningún error.")
}

func TestGeneracionAutomagica_Error_DuracionNoAlcanzada(t *testing.T) {
	// Crear una lista de ejercicios y agregar 2 ejercicios
	listaEjercicios := NewListaDeEjercicios()
	err := listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripción de Flexiones de brazos", 20, 50, []string{"fuerza"}, []int{10}, "principiante")
	if err != nil {
		t.Errorf("Error al agregar ejercicio: %v", err)
	}
	err = listaEjercicios.AgregarEjercicio("Sentadillas", "Descripción de Sentadillas", 25, 50, []string{"fuerza"}, []int{10}, "principiante")
	if err != nil {
		t.Errorf("Error al agregar ejercicio: %v", err)
	}
	// Crear una lista de rutinas
	listaDeRutinas := NewListaDeRutinas()
	// Llamar a la función GeneracionAutomagica con una duración total mayor que la duración de los ejercicios disponibles
	_, err = listaDeRutinas.GeneracionAutomagica("Rutina3", 50, "fuerza", "principiante", listaEjercicios)
	assert.Error(t, err, "Se esperaba un error debido a que la duración total no puede ser alcanzada con los ejercicios disponibles, pero no se recibió ningún error.")
}

// Test Generación Automágica2 de Rutinas
func TestGeneracionAutomagica2_Exito(t *testing.T) {
	// Crear una lista de ejercicios y agregar un ejercicio
	listaEjercicios := NewListaDeEjercicios()
	err := listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripción de Flexiones de brazos", 30, 50, []string{"fuerza"}, []int{10}, "principiante")
	if err != nil {
		t.Errorf("Error al agregar ejercicio: %v", err)
	}

	// Crear una lista de rutinas
	listaDeRutinas := NewListaDeRutinas()

	// Llamar a la función GeneracionAutomagica2
	rutina, err := listaDeRutinas.GeneracionAutomagica2("Rutina1", 50, listaEjercicios)

	// Verificar que no haya ocurrido un error
	assert.Nil(t, err, "Se esperaba que no ocurriera ningún error")

	// Verificar que la cantidad de calorías de la rutina sea igual a la de calorías objetivo
	assert.Equal(t, 50, rutina.caloriasQuemadasTotales, "La duración no es igual a la duración total de los ejercicios disponibles")
}

func TestGeneracionAutomagica2_Error_CaloriasInsuficientes(t *testing.T) {
	// Crear una lista de ejercicios y agregar un ejercicio
	listaEjercicios := NewListaDeEjercicios()
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripción de Flexiones de brazos", 30, 50, []string{"fuerza"}, []int{10}, "principiante")

	// Crear una lista de rutinas
	listaDeRutinas := NewListaDeRutinas()

	// Llamar a la función GeneracionAutomagica2 con calorías objetivo mayores que las disponibles
	_, err := listaDeRutinas.GeneracionAutomagica2("Rutina1", 100, listaEjercicios)

	// Verificar que se haya recibido un error
	assert.Error(t, err, "Se esperaba un error debido a las calorías insuficientes, pero no se recibió ninguno")
}

// Test Generación Automágica3 de Rutinas

func TestGeneracionAutomagica3_Exito(t *testing.T) {
	// Crear una lista de ejercicios y agregar un ejercicio
	listaEjercicios := NewListaDeEjercicios()
	err := listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripción de Flexiones de brazos", 30, 50, []string{"fuerza"}, []int{10}, "principiante")
	if err != nil {
		t.Errorf("Error al agregar ejercicio: %v", err)
	}

	// Crear una lista de rutinas
	listaDeRutinas := NewListaDeRutinas()

	// Llamar a la función GeneracionAutomagica3
	rutina, err := listaDeRutinas.GeneracionAutomagica3("Rutina1", 50, "fuerza", listaEjercicios)

	// Verificar que no haya ocurrido un error
	assert.Nil(t, err, "Se esperaba que no ocurriera ningún error")

	// Verificar que la cantidad de calorías de la rutina sea igual a la de calorías objetivo
	assert.Equal(t, 50, rutina.caloriasQuemadasTotales, "La duración no es igual a la duración total de los ejercicios disponibles")
}

// ///// Automagica3 v2
func TestGeneracionAutomagica3v2(t *testing.T) {
	// Crear una lista de ejercicios y agregar varios ejercicios
	listaEjercicios := NewListaDeEjercicios()
	listaEjercicios.AgregarEjercicio("Flexiones de brazos", "Descripción de Flexiones de brazos", 30, 50, []string{"fuerza"}, []int{20}, "principiante")
	listaEjercicios.AgregarEjercicio("Sentadillas", "Descripción de sentadillas", 30, 60, []string{"fuerza"}, []int{30}, "principiante")
	listaEjercicios.AgregarEjercicio("Estocadas", "Descripción de estocadas", 30, 70, []string{"fuerza"}, []int{40}, "principiante")

	// Crear una lista de rutinas
	listaDeRutinas := NewListaDeRutinas()

	// Llamar a la función GeneracionAutomagica3v2
	rutina, err := listaDeRutinas.GeneracionAutomagica3v2("Rutina1", 60, "fuerza", listaEjercicios)
	if err != nil {
		t.Fatalf("Error al generar la rutina automática: %v", err)
	}

	// Verificar la duración de la rutina
	assert.Equal(t, 60, rutina.duracion, "La duración de la rutina no coincide")

	// Verificar que la rutina contenga los ejercicios esperados
	assert.Equal(t, 2, len(rutina.ejerciciosTotales), "La rutina debería contener al menos un ejercicio")

	// Preparar los nombres esperados de los ejercicios
	nombresEsperados := []string{"Sentadillas", "Estocadas"}

	// Verificar que los nombres de los ejercicios generados coincidan con los esperados
	for _, ejercicio := range rutina.ejerciciosTotales {
		encontrado := false
		for _, nombre := range nombresEsperados {
			if ejercicio.nombre == nombre {
				encontrado = true
				break
			}
		}
		assert.True(t, encontrado, "El ejercicio generado no coincide con los esperados")
	}
}
