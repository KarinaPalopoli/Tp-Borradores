package rutinaDeEjercicios

import (
	"reflect"
	"testing"
)

func TestAgregarRutina(t *testing.T) {
	lista := NewListaDeRutinas()

	// Caso de prueba 1: Agregar una rutina válida
	ejercicios1 := []*Ejercicio{
		{nombre: "Flexiones", tiempo: 30, calorias: 100, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
		{nombre: "Sentadillas", tiempo: 45, calorias: 150, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
	}
	err := lista.AgregarRutina("Rutina1", ejercicios1)
	if err != nil {
		t.Errorf("Error al agregar la rutina 1: %v", err)
	}

	// Caso de prueba 2: Intentar agregar una rutina con nombre duplicado
	err = lista.AgregarRutina("Rutina1", ejercicios1)
	if err == nil {
		t.Errorf("Se esperaba un error al intentar agregar una rutina con nombre duplicado, pero no se recibió")
	}
}

func TestEliminarRutina(t *testing.T) {
	lista := NewListaDeRutinas()

	ejercicios1 := []*Ejercicio{
		{nombre: "Flexiones", tiempo: 30, calorias: 100, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
		{nombre: "Sentadillas", tiempo: 45, calorias: 150, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
	}
	err := lista.AgregarRutina("Rutina1", ejercicios1)
	if err != nil {
		t.Errorf("Error al agregar la rutina 1: %v", err)
	}

	// Caso de prueba 1: Eliminar una rutina existente
	err = lista.BorrarRutina("Rutina1")
	if err != nil {
		t.Errorf("Error al eliminar la rutina 1: %v", err)
	}

	// Caso de prueba 2: Intentar eliminar una rutina que no existe
	err = lista.BorrarRutina("Rutina1")
	if err == nil {
		t.Errorf("Se esperaba un error al intentar eliminar una rutina que no existe, pero no se recibió")
	}
}

func TestConsultarRutina(t *testing.T) {
	lista := NewListaDeRutinas()

	ejercicios1 := []*Ejercicio{
		{nombre: "Flexiones", tiempo: 30, calorias: 100, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
		{nombre: "Sentadillas", tiempo: 45, calorias: 150, tipoDeEjercicio: []string{"fuerza"}, dificultad: "principiante"},
	}
	err := lista.AgregarRutina("Rutina1", ejercicios1)
	if err != nil {
		t.Errorf("Error al agregar la rutina 1: %v", err)
	}

	// Caso de prueba 1: Consultar una rutina existente
	rutina, err := lista.ConsultarRutina("Rutina1")
	if err != nil {
		t.Errorf("Error al consultar la rutina 1: %v", err)
	}
	if !reflect.DeepEqual(rutina.ejerciciosTotales, ejercicios1) {
		t.Errorf("Los ejercicios de la rutina 1 no coinciden con los esperados")
	}

	// Caso de prueba 2: Intentar consultar una rutina que no existe
	_, err = lista.ConsultarRutina("Rutina2")
	if err == nil {
		t.Errorf("Se esperaba un error al intentar consultar una rutina que no existe, pero no se recibió")
	}
}
func TestCalcularDuracion(t *testing.T) {
	// Caso de prueba 1: Lista de ejercicios vacía
	ejercicios := []*Ejercicio{}
	duracion := calcularDuracion(ejercicios)
	if duracion != 0 {
		t.Errorf("Error en el caso de prueba 1. Se esperaba una duración de 0 pero se obtuvo %d", duracion)
	}

	// Caso de prueba 2: Lista de ejercicios con elementos
	ejercicios = []*Ejercicio{
		{nombre: "Flexiones", tiempo: 30},
		{nombre: "Sentadillas", tiempo: 45},
		{nombre: "Abdominales", tiempo: 20},
	}
	duracionEsperada := 30 + 45 + 20 // Duración total esperada es la suma de los tiempos de cada ejercicio
	duracion = calcularDuracion(ejercicios)
	if duracion != duracionEsperada {
		t.Errorf("Error en el caso de prueba 2. Se esperaba una duración de %d pero se obtuvo %d", duracionEsperada, duracion)
	}
}