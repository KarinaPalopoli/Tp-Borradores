package tp2024grupo777

import (
	"github.com/untref-ayp2/data-structures"

)

// Definición de  estructura 
type Rutina[T any] struct {
	id int
	nombre string
	ejerciciosIncluidos []Ejercicio
	}



type TipoDeEjercicio string
 const (
	cardio    tipoDeEjercicio = "cardio"
    tonicidad tipoDeEjercicio = "tonicidad"
   	balance   tipoDeEjercicio = "balance"
)

type NivelDeDificultad string
 const (
	principiante    nivelDeDificultad = "principiante"
    intermedio nivelDeDificultad = "intermedio"
	avanzado   nivelDeDificultad = "avanzado"
)

type Ejercicio[T any] struct {  //cada ejercicio tendra su etiqueta para definir
	// tipoDeEjercicio y nivelDeDificultad
	nombre string
	descripcion string
	duracion int
	nivelDeDificultad int
	caloriasQuemadas int
	tipoDeEjercicio TipoDeEjercicio  //UN EJERCICIO PUEDE TENER MAS DE UNA CATEGORIA
	puntosPorTipoDeEjercicio map[TipoDeEjercicio]int  // puee tener puntaje en más de una categoria
	nivelDificultad NivelDeDificultad

}
	
//las etiquetas seran un conj de palabras clave que permitiran clasificar los ejercicios
// Ej: Sentadilla pertenece al grupo "tonicidad" y "cardio".  PERO LAS DIFICULTADES SSON UNICASD, 
// SENTADILLA SERA DIFICULTAD ,MEDIA SOLAMENTE	

// Definición de la estructura Resumen de Rutina
type Resumen struct {
	caloriasQuemadasTotales int 
	puntosCategoria map[string]int
 } // Mapa donde se almacenan 

// los puntos por dimensión (por ejemplo, "Tone", "Cardio", "Balance"...)



// Función para calcular la duración total de una rutina en minutos
func (r *Rutina) duracionTotal() int {
    totalDuracion := 0
    for _, ejercicio := range r.ejerciciosIncluidos {
        //  cálculo relacionado con la duración de cada ejercicio
       
        totalDuracion += ejercicio.nivelDeDificultad
    }
    return totalDuracion
}

// Función para calcular las calorías quemadas totales de una rutina
func (r *Rutina) caloriasQuemadasTotales() int {
    totalCalorias := 0
    for _, ejercicio := range r. {
        // sumar las calorías quemadas de cada ejercicio a las calorías totales
        totalCalorias += ejercicio.caloriasQuemadas
    }
    return totalCalorias
}





	


	


// Crear instancia de Rutina
rutina1 := Rutina{
    Nombre: "Rutina matutina",
    DuracionTotal: 60,
    EjerciciosIncluidos: []Ejercicios{ejercicio1},
}




// Crear instancia de Resumen
resumen1 := Resumen{
   CaloriasQuemadasTotales: 200,
   PuntosDimensión :map[string]int{
 "Tone": 10,
 "Cardio": 20,
 "Balance": 5,
}
}