package rutinaDeEjercicios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAgregarEjercicio(t *testing.T) {
    lista := NewListaDeEjercicios()
    
    // Caso de prueba exitoso
    err := lista.AgregarEjercicio("Flexiones de brazos", "Se realiza estando en posición inclinada, recostado hacia abajo, levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo", 10, 100, []string{"fuerza"}, []int{50}, "principiante")
    assert.NoError(t, err, "Se esperaba que no haya error al agregar el ejercicio")
    
    // Caso de prueba donde el ejercicio ya existe y da error
    err = lista.AgregarEjercicio("Flexiones de brazos", "Se realiza estando en posición inclinada, recostado hacia abajo, levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo", 10, 100, []string{"fuerza"}, []int{50}, "principiante")
    assert.Error(t, err, "Se esperaba un error al agregar un ejercicio que ya existe")
}

func TestBorrarEjercicio(t *testing.T) {
    lista := NewListaDeEjercicios()
    
    // Agregar ejercicio para poder borrarlo después
    lista.AgregarEjercicio("Flexiones de brazos", "Se realiza estando en posición inclinada, recostado hacia abajo, levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo", 10, 100, []string{"fuerza"}, []int{50}, "principiante")
    
    // Caso de prueba exitoso
    err := lista.BorrarEjercicio("Flexiones de brazos")
    assert.NoError(t, err, "Se esperaba que no haya error al borrar el ejercicio")
    
    // Caso de prueba donde el ejercicio no existe
    err = lista.BorrarEjercicio("Flexiones de brazos")
    assert.Error(t, err, "Se esperaba un error al intentar borrar un ejercicio que ya no existe")
}

func ConsultarEjercicioPorNombre(t *testing.T) {
    lista := NewListaDeEjercicios()
    
    // Agregar ejercicio para poder después consultarlo
    ejercicio := Ejercicio{
        nombre: "Flexiones de brazos",
        descripcion: "Se realiza estando en posición inclinada, recostado hacia abajo, levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo",
        tiempo: 10,
        calorias: 100,
        tipoDeEjercicio: []string{"fuerza"},
        puntosPorTipoDeEjercicio: []int{50},
        dificultad: "principiante",
    }
    lista.listaDeEjercicios["Flexiones de brazos"] = &ejercicio
    
    // Caso de prueba exitoso
	valor, err := lista.ConsultarEjercicioPorNombre("Flexiones de brazos")
    assert.NoError(t, err, "Se esperaba que no haya error al consultar el ejercicio")
    assert.Equal(t, ejercicio, *valor, "Los ejercicios deberían ser iguales")
	assert.Equal(t, ejercicio.dificultad, valor.dificultad, "Los ejercicios deberían ser iguales")
	assert.Equal(t, ejercicio.puntosPorTipoDeEjercicio, valor.puntosPorTipoDeEjercicio, "Los ejercicios deberían ser iguales")
	assert.Equal(t, ejercicio.descripcion, valor.descripcion, "Los ejercicios deberían ser iguales")
	assert.Equal(t, ejercicio.tiempo, valor.tiempo, "Los ejercicios deberían ser iguales")

    
    // Caso de prueba donde el ejercicio no existe
    _, err = lista.ConsultarEjercicioPorNombre("Sentadillas")
    assert.Error(t, err, "Se esperaba un error al consultar un ejercicio que no existe")
}

func TestModificarEjercicio(t *testing.T) {
    lista := NewListaDeEjercicios()
    
    // Agregar ejercicio para poder despues modificarlo
    lista.AgregarEjercicio("Flexiones de brazos", "Se realiza estando en posición inclinada, recostado hacia abajo, levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo", 10, 100, []string{"fuerza"}, []int{50}, "principiante")
    
    // Caso de prueba exitoso
    err := lista.ModificarEjercicio("Flexiones de brazos", "Nueva descripción de flexiones de brazos", 15, 150, []string{"fuerza"}, []int{60}, "intermedio")
    assert.NoError(t, err, "Se esperaba que no haya error al modificar el ejercicio")
    
    // Caso de prueba donde el ejercicio no existe
    err = lista.ModificarEjercicio("Sentadillas", "flexionar las rodillas y bajar el cuerpo manteniendo la verticalidad, para luego regresar a una posición erguida", 20, 200, []string{"fuerza"}, []int{70}, "avanzado")
    assert.Error(t, err, "Se esperaba un error al modificar un ejercicio que no existe")
}

func TestListarEjercicios(t *testing.T) {
    lista := NewListaDeEjercicios()
    
    // Agregar primer ejercicio
    ejercicio1 := Ejercicio{
        nombre: "Flexiones de brazos",
        descripcion: "Se realiza estando en posición inclinada, recostado hacia abajo, levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo",
        tiempo: 10,
        calorias: 100,
        tipoDeEjercicio: []string{"fuerza"},
        puntosPorTipoDeEjercicio: []int{50},
        dificultad: "principiante",
    }
    lista.listaDeEjercicios["Flexiones de brazos"] = &ejercicio1

	ejercicio2 := Ejercicio{
        nombre: "Sentadillas",
        descripcion: "Flexionar las rodillas y bajar el cuerpo manteniendo la verticalidad, para luego regresar a una posición erguida",
        tiempo: 7,
        calorias: 90,
        tipoDeEjercicio: []string{"fuerza"},
        puntosPorTipoDeEjercicio: []int{60},
        dificultad: "principiante",
    }
    lista.listaDeEjercicios["Sentadillas"] = &ejercicio2
   
    // Caso de prueba exitoso
	ejercicios,_:= lista.ListarEjercicios()
	
	// Verificar la longitud del slice devuelto
    assert.Equal(t, 2, len(ejercicios), "Se espera que haya 2 ejercicios en la listaDeEjercicios")
    assert.Equal(t, &ejercicio1, ejercicios[0], "Los ejercicios deberían ser iguales")
	assert.Equal(t, &ejercicio2, ejercicios[1], "Los ejercicios deberían ser iguales")

	// Verificar que los ejercicios de la lista sean los correctos
	assert.Contains(t, ejercicios, &ejercicio1, "El ejercicio1 que es Flexiones de brazos, debería estar en la lista")
	assert.Contains(t, ejercicios, &ejercicio2, "El ejercicio1 que es Sentadillas, debería estar en la lista")

}

func TestFiltrarEjercicios(t *testing.T){
    lista:= NewListaDeEjercicios()
    // Agregamos ejercicios para después filtrar
    ejercicio1 := Ejercicio{
        nombre: "Flexiones de brazos",
        descripcion: "Se realiza estando en posición inclinada, recostado hacia abajo, levantando el cuerpo únicamente con los brazos y bajando de nuevo al suelo",
        tiempo: 10,
        calorias: 100,
        tipoDeEjercicio: []string{"fuerza"},
        puntosPorTipoDeEjercicio: []int{50},
        dificultad: "principiante",
    }
    lista.listaDeEjercicios["Flexiones de brazos"] = &ejercicio1

	ejercicio2 := Ejercicio{
        nombre: "Sentadillas",
        descripcion: "Flexionar las rodillas y bajar el cuerpo manteniendo la verticalidad, para luego regresar a una posición erguida",
        tiempo: 7,
        calorias: 90,
        tipoDeEjercicio: []string{"fuerza"},
        puntosPorTipoDeEjercicio: []int{60},
        dificultad: "principiante",
    }
    lista.listaDeEjercicios["Sentadillas"] = &ejercicio2
    ejercicio3 := Ejercicio{
        nombre: "Estocadas",
        descripcion: "Separar los pies, dar un paso hacia adelante. Mantener el peso en la pierna que está al frente. Cambiar de pierna y repetir",
        tiempo: 12,
        calorias: 20,
        tipoDeEjercicio: []string{"balance"},
        puntosPorTipoDeEjercicio: []int{30},
        dificultad: "intermedio",
    }
    lista.listaDeEjercicios["Estocadas"] = &ejercicio3

    // Caso de prueba donde se filtra por tipo de ejercicio
    ejerciciosFiltrados,error:= lista.FiltrarEjercicios("balance", "", 0)
    assert.Equal(t, &ejercicio3, ejerciciosFiltrados[0], "Los ejercicios deberían ser iguales")
    assert.NoError(t, error) // No debería haber error

    // Caso de prueba donde se filtra por mínimo de calorías
    ejerciciosFiltrados,_= lista.FiltrarEjercicios("", "", 90) // Debería traer un slice con 2 ejercicios
    assert.Equal(t,2,len(ejerciciosFiltrados))

    // Caso de prueba donde se filtra por dificultad
    ejerciciosFiltrados,_= lista.FiltrarEjercicios("", "intermedio", 0)
    assert.Equal(t, &ejercicio3, ejerciciosFiltrados[0], "Los ejercicios deberían ser iguales")

    // Caso de prueba donde no se pasan los filtros vacíos, debería traer un slice con 3 elementos
    ejerciciosFiltrados,_= lista.FiltrarEjercicios("", "", 0)
    assert.Equal(t,3,len(ejerciciosFiltrados))
    
    // Caso de prueba donde no se encuentra ejercicio que cumpla los criterios de filtro. Debería dar error
    ejerciciosFiltrados,error= lista.FiltrarEjercicios("", "", 200)
    assert.Error(t, error) // Debería haber error
}