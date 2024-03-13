//Participantes :
//					Andrea Iasbel Ramirez Solarte -  Estefanía Valencia Vallejo

package main

import (
	"github.com/estefa0942/proyectoEstudioDistribucion/calculadora"
	"github.com/estefa0942/proyectoEstudioDistribucion/gestorArchivos"
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Experimentación ST vs MT")

	instancia := "./InstanciasConcurrencia/01_5.input"
	// instancia := "./InstanciasConcurrencia/02_20.input"
	// instancia := "./InstanciasConcurrencia/03_40.input"
	// instancia := "./InstanciasConcurrencia/04_100.input"
	// instancia := "./InstanciasConcurrencia/05_200.input"
	// instancia := "./InstanciasConcurrencia/06_500.input"
	// instancia := "./InstanciasConcurrencia/07_1000.input"
	// instancia := "./InstanciasConcurrencia/08_2500.input"
	// instancia := "./InstanciasConcurrencia/09_10000.input"

	//Versión SingleThreaded
	// fmt.Println("-> Ejecutando versión de un solo hilo")
	// coleccionOperaciones := gestorArchivos.CargarArchivo(instancia)
	// tiempoInicial := time.Now()
	// for i := 0; i < len(coleccionOperaciones); i++ {
	// 	coleccionOperaciones[i].Operar()
	// }

	// var suma float32
	// suma = 0
	// for i := 0; i < len(coleccionOperaciones); i++ {
	// 	suma += coleccionOperaciones[i].Resultado
	// }
	// fmt.Println("Resultado: ", suma)

	// tiempoFinal := time.Since(tiempoInicial)
	// fmt.Println("Tiempo transcurrido: ", tiempoFinal)

	//Versión MultiThreaded - Data Race
	// fmt.Println("-> Ejecutando versión de una variable")

	// var suma float32
	// var wg sync.WaitGroup
	// for i := 0; i < len(coleccionOperaciones); i++ {
	// 	wg.Add(1)
	// 	go func(x int) {
	// 		coleccionOperaciones[x].Operar()
	// 		suma += coleccionOperaciones[x].Resultado
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

	// fmt.Println("Resultado: ", suma)
	// tiempoFinal := time.Since(tiempoInicial)
	// fmt.Println("Tiempo transcurrido: ", tiempoFinal.Seconds())

	//Versión MultiThreaded - Colección reservada
	// fmt.Println("-> Ejecutando versión de colección reservada")
	// coleccionOperaciones := gestorArchivos.CargarArchivo(instancia)
	// tiempoInicial := time.Now()

	// suma := make([]float32, len(coleccionOperaciones))

	// var sumas float32
	// var wg sync.WaitGroup

	// for i := 0; i < len(coleccionOperaciones); i++ {
	// 	wg.Add(1)
	// 	go func(op calculadora.Operacion, ii int) {
	// 		op.Operar()
	// 		suma[ii] += op.Resultado
	// 		wg.Done()
	// 	}(coleccionOperaciones[i], i)

	// }
	// wg.Wait()
	// for j, _ := range suma {
	// 	sumas += suma[j]
	// }

	// fmt.Println("Resultado: ", sumas)
	// tiempoFinal := time.Since(tiempoInicial)
	// fmt.Println("Tiempo transcurrido: ", tiempoFinal.Seconds())

	//Versión MultiThreaded - Semáforos
	fmt.Println("-> Ejecutando versión de colección reservada")
	coleccionOperaciones := gestorArchivos.CargarArchivo(instancia)
	tiempoInicial := time.Now()

	var suma float32
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < len(coleccionOperaciones); i++ {
		wg.Add(1)
		go func(op calculadora.Operacion) {
			defer wg.Done()
			op.Operar()
			mutex.Lock()
			suma += op.Resultado
			mutex.Unlock()
		}(coleccionOperaciones[i])
	}

	wg.Wait()

	fmt.Println("Resultado: ", suma)
	tiempoFinal := time.Since(tiempoInicial)
	fmt.Println("Tiempo transcurrido: ", tiempoFinal.Seconds())

}
