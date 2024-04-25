package rutinaDeEjercicios

import (
	"testing"
)

// VERIFICA EL BORRADO DE UNA RUTINA ACTUAL
func TestBorrarRutinaExistente(t *testing.T) {
	rutinas := make(map[string]*Rutina)
	lista := ListaDeRutinas{listaDeRutinas: rutinas}
	rutinas["Rutina de Lunes"] = &Rutina{nombre: "Rutina de Lunes"}

	err := lista.BorrarRutina("Rutina de Lunes")
	if err != nil {
		t.Errorf("Error al borrar una rutina existente: %s", err)
	}

	if _, ok := rutinas["Rutina de Lunes"]; ok {
		t.Error("La rutina no fue borrada correctamente")
	}
}

// VERIFICA QUE NO EXITE LA RUTINA
func TestBorrarRutinaNoExistente(t *testing.T) {
	rutinas := make(map[string]*Rutina)
	lista := ListaDeRutinas{listaDeRutinas: rutinas}

	err := lista.BorrarRutina("Rutina de martes")
	if err == nil || err.Error() != "la rutina no existe" {
		t.Errorf("Expected error 'la rutina no existe', got %v", err)
	}
}

// VERIFICA EN EL CASO DE AGREGAR UNA RUTINA DUPICADA

func TestAgregarEjercicioARutinaExistente(t *testing.T) {
	lista := NewListaDeRutinas()
	ejercicio := &Ejercicio{nombre: "abdominales", tiempo: 5, calorias: 50, tipoDeEjercicio: []string{"CARDIO"}, dificultad: "INTERMEDIA"}
	lista.listaDeRutinas["Rutina de Martes"] = &Rutina{
		nombre:            "Rutina de Martes",
		ejerciciosTotales: []*Ejercicio{ejercicio},
	}

	err := lista.AgregarEjercicioARutina("Rutina de Martes", ejercicio)
	if err == nil || err.Error() != "el ejercicio ya está dentro de la rutina" {
		t.Error("Expected Error")
	}
}

// VERIFICA AGREGAR RUTINA SIN EJERCICIO
func TestAgregarEjercicioARutinaInexistente(t *testing.T) {
	lista := NewListaDeRutinas()
	ejercicio := &Ejercicio{nombre: "sentadillas", tiempo: 10, calorias: 130, tipoDeEjercicio: []string{"TONICIDAD"}, dificultad: "INTERMEDIO"}

	err := lista.AgregarEjercicioARutina("Rutina Inexistente", ejercicio)
	if err == nil || err.Error() != "la rutina no existe" {
		t.Errorf("Expected 'la rutina no existe' error, got: %v", err)
	}
}