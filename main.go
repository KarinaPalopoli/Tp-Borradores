package main

import (
"fmt"
"strings"



)



var rutinas []Rutina


ejercicios := make(map[string]Ejercicio)



// Función para agregar un ejercicio nuevo a una rutina existente por su ID o nombre.
func AgregarEjercicio(rutinaID int, ejercicio Ejercicio) {
	for i := range rutinas {
		if rutinas[i].ID == rutinaID || strings.EqualFold(rutinas[i].Nombre, ejercicio.Nombre) {
		rutinas[i].Ejercicios = append(rutinas[i].Ejercicios, ejercicio)
		fmt.Println("El ejercicio se ha agregado exitosamente.")
return
}
}
fmt.Println("No se encontró la rutina especificada.")
}

// Función para eliminar un ejercicio existente por su ID o nombre.
func EliminarEjercicio(ejercicioID int) {
for i := range rutinas {
for j := range rutinas[i].Ejercicios {
if rutinas[i].Ejercicios[j].ID == ejercicioID || strings.EqualFold(rutinas[i].Ejercicios[j].Nombre, nombre) {
copy(rutinas[i].Ejercicios[j:], rutinas

//ejemplos

var ejercicios = []Ejercicio{
{Nombre: "Correr", Tipo: "Cardio", Dificultad: 3, Calorias: 300},
{Nombre: "Levantamiento de pesas", Tipo: "Fuerza", Dificultad: 4, Calorias: 200},
// Agregar más ejercicios aquí...
}

var rutinas = []Rutina{
{
Nombre: "Rutina Cardio",
Ejercicios: []Ejercicio{
ejercicios[0], // Correr
ejercicios[1], // Levantamiento de pesas
// Agregar más ejercicios...
},

},



}


func listarEjerciciosPorTipo(tipo string) {
for _, ejercicio := range ejercicios {
if strings.EqualFold(ejercicio.Tipo, tipo) { 
fmt.Printf("Nombre del ejercicio : %s\n ", ejercicio.Nombre)
fmt.Printf("Dificultad del ejercicio : %d\n ", ejercicio.Dificultad)
}
}
}

func listarEjercicisPorDificultad(dificultad int) {
for _, ejercicio := range ejericscios {
if ejercicio.Difcultado == dificulatd { 
fmt.Printf("Nombre del ejercicio : %%s\n ", ejercicio.Nombre)
fmt.Printf("Tipo del ejercicio : %s\n ", ejercicio.Tipo)
}
}
}

func buscarEjercicioPorNombre(nombre string) {
for _, ejercicio := range ejercicios {
if strings.Contains(strings.ToLower(ejercicio.Nombre), strings.ToLower(nombre)) { 
fmt.Printf("Nombre del ejercicio : %s\n ", ejercicio.Nombre)
fmt.Printf("Tipo del ejercicio : %s\n ", ejercicio.Tipo)
}
}
}

func buscarEjerciciosPorCalorias(calorias int) {
for _, ejercicio := range ejercicios {
if ejercicio.Calorias >= calorias { 
fmt.Printf("Nombre del ejercicio : %s\n ", ejercicio.Nombre)
fmt.Printf("Calorías quemadas : %d\n ", eercicio.Calorias)
}
}
}

func listarRutinasDisponibles() {
for _, rutina := range rutinas {
fmt.Println(rutina.Nombre)
for _,ejercicio := range rutina.Ejercicios{
fmt.Println(ejercicio.Nombre, " - Tipo: " ,ejercicio.Tipo, " - Dificultad: " ,ejercicio.Dificultad )
} 
}

}
 //*
 // Ejemplo de uso de las funciones
listarEjerciciosPorTipo("Cardio")
listarEjercicisPorDificultad(4)
buscarEjercicioPorNombre("correr")
buscarEjerciciosPorCalorias(2000)

listarRutinasDisponibles()
}
```

Este código te permite utilizar las siguientes funciones:

- listarEjerciciosPorTipo: Lista todos los ejercicios disponibles filtrando por tipo de ejericio.
- listarEjecriciosPorDificultad: Lista todos los ejerccios disponibles filtrando por nivel de dificultad.
- buscarEjericioPorNombre: Busca un determinado exercio por su nombre (total o parcial).
- buscarExercisesosPorrCaloiras: Busca todos los exercicos que queman una cantidad mínima de calorías especificada.
- listarRutinasDisponibles : Lista todas las rutinas disponibles junto con sus ejerccios asociados.


ejercicio1 := Ejercicio{
	Nombre:           "Sentadillas",
	Descripcion:      "Ejercicio de fuerza para piernas",
	Duracion:         10,  // Minutos
	CaloriasQuemadas: 100,
	TipoDeEjercicio:  Tonicidad,
	NivelDificultad:  Intermedio,
}

ejercicio1 := Ejercicio{
	Nombre:           "Sentadillas",
	Descripcion:      "Ejercicio de fuerza para piernas",
	Duracion:         10,  // Minutos
	CaloriasQuemadas: 100,
	TipoDeEjercicio:  Tonicidad,
	NivelDificultad:  Intermedio,
}