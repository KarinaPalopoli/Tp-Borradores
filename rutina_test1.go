package rutinaDeEjercicios

import (
    "testing"

    "github.com/stretchr/testify/assert"
   
)

func TestCalcularDuracion(t *testing.T) {
    // Prueba con Lista de ejercicios vacía
    ejercicios := []*rutinaDeEjercicios.Ejercicio{}
    assert.Equal(t, 0, rutinaDeEjercicios.CalcularDuracion(ejercicios), "La duración debe ser 0 para una lista vacía")

    // Prueba con Lista de ejercicios con elementos
    ejercicios = []*rutinaDeEjercicios.Ejercicio{
        {Nombre: "Flexiones", Tiempo: 30},
        {Nombre: "Sentadillas", Tiempo: 45},
        {Nombre: "Abdominales", Tiempo: 20},
    }
    duracionEsperada := 30 + 45 + 20 
    assert.Equal(t, duracionEsperada, rutinaDeEjercicios.CalcularDuracion(ejercicios), "La duración no coincide con la esperada")
 }
 
 func TestAgregarRutina(t *testing.T) {
    lista := rutinaDeEjercicios.NewListaDeRutinas()

    // prueba con una rutina válida
    ejercicios1 := []*rutinaDeEjercicios.Ejercicio{
        {Nombre: "Flexiones", Tiempo: 30},
        {Nombre: "Sentadillas", Tiempo: 45},
    }
    err := lista.AgregarRutina("Rutina1", ejercicios1)
    assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

    // prueba con nombre de rutina duplicado
    err = lista.AgregarRutina("Rutina1", ejercicios1)
    assert.Error(t, err, "Se esperaba un error al intentar agregar una rutina con nombre duplicado")
    assert.EqualError(t, err, "la rutina ya existe")

    // Agregando rutina sin ejercicios
    err = lista.AgregarRutina("Rutina2", []*rutinaDeEjercicios.Ejercicio{})
    assert.Error(t, err, "Se esperaba un error al intentar agregar una rutina sin ejercicios")
    assert.EqualError(t, err, "una rutina debe contener al menos 1 ejercicio")
}

func TestBorrarRutina(t *testing.T) {
    lista := rutinaDeEjercicios.NewListaDeRutinas()
	//se crea una rutina con 1 ejercicio
    ejercicios1 := []*rutinaDeEjercicios.Ejercicio{
        {Nombre: "Flexiones", Tiempo: 30},
        {Nombre: "Sentadillas", Tiempo: 45},
    }
    err := lista.AgregarRutina("Rutina1", ejercicios1)
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
    lista := rutinaDeEjercicios.NewListaDeRutinas()

    ejercicios1 := []*rutinaDeEjercicios.Ejercicio{
        {Nombre: "Flexiones", Tiempo: 30},
        {Nombre: "Sentadillas", Tiempo: 45},
    }
    err := lista.AgregarRutina("Rutina1", ejercicios1)
    assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

    // consultar una rutina existente
    rutina, err := lista.ConsultarRutina("Rutina1")
    assert.NoError(t, err, "No se esperaba un error al consultar la rutina 1")
    assert.Equal(t, "Rutina1", rutina.Nombre, "El nombre de la rutina no coincide")

    // intentar consultar una rutina inexistente
    _, err = lista.ConsultarRutina("Rutina2")
    assert.Error(t, err, "Se esperaba un error al intentar consultar una rutina que no existe")
    assert.EqualError(t, err, "la rutina no existe")
}

func TestModificarRutina(t *testing.T) {
    lista := rutinaDeEjercicios.NewListaDeRutinas()

    ejercicios1 := []*rutinaDeEjercicios.Ejercicio{
        {Nombre: "Flexiones", Tiempo: 30},
        {Nombre: "Sentadillas", Tiempo: 45},
    }
    err := lista.AgregarRutina("Rutina1", ejercicios1)
    assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

    // modificando rutina existente
    nuevosEjercicios := []*rutinaDeEjercicios.Ejercicio{
        {Nombre: "Abdominales", Tiempo: 20},
    }
    err = lista.ModificarRutina("Rutina1", nuevosEjercicios)
    assert.NoError(t, err, "No se esperaba un error al modificar la rutina 1")

    // intentar modificar una rutina inexistente
    err = lista.ModificarRutina("Rutina2", nuevosEjercicios)
    assert.Error(t, err, "Se esperaba un error al intentar modificar una rutina que no existe")
    assert.EqualError(t, err, "la rutina no existe")
}

func TestListarRutinas(t *testing.T) {
    lista := rutinaDeEjercicios.NewListaDeRutinas()

    // listar sin rutinas
    rutinas, err := lista.ListarRutinas()
    assert.Nil(t, rutinas, "Se esperaba nil al listar rutinas cuando no hay ninguna")
    assert.Error(t, err, "Se esperaba un error al listar rutinas cuando no hay ninguna")
    assert.EqualError(t, err, "no hay ninguna rutina para listar")

    ejercicios1 := []*rutinaDeEjercicios.Ejercicio{
        {Nombre: "Flexiones", Tiempo: 30},
        {Nombre: "Sentadillas", Tiempo: 45},
    }
    err = lista.AgregarRutina("Rutina1", ejercicios1)
    assert.NoError(t, err, "No se esperaba un error al agregar la rutina 1")

    // listar con al menos una rutina
    rutinas, err = lista.ListarRutinas()
    assert.NotNil(t, rutinas, "Se esperaba una lista de rutinas al listar rutinas cuando hay al menos una")
    assert.NoError(t, err, "No se esperaba un error al listar rutinas cuando hay al menos una")
    assert.Equal(t, 1, len(rutinas), "Se esperaba una lista de rutinas con un elemento")
}
