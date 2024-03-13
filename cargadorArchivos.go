package gestorArchivos

import (
	"bufio"
	"github.com/estefa0942/proyectoEstudioDistribucion/calculadora"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func CargarArchivo(nombreArchivo string) []calculadora.Operacion {

	//Preparar slice de salida
	var coleccionOperaciones []calculadora.Operacion

	// Abre el archivo para leer
	file, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Lee cada línea del archivo
	for scanner.Scan() {
		linea := scanner.Text()
		campos := strings.Split(linea, " ")

		// Convierte los valores a los tipos adecuados
		a, _ := strconv.ParseFloat(campos[0], 32)
		b, _ := strconv.ParseFloat(campos[1], 32)
		operador := campos[2]
		cargaComputacional, _ := strconv.Atoi(campos[3])

		// Crea la estructura Operacion
		operacion := calculadora.Operacion{
			A:                  float32(a),
			B:                  float32(b),
			Operador:           operador,
			CargaComputacional: cargaComputacional,
		}

		coleccionOperaciones = append(coleccionOperaciones, operacion)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return coleccionOperaciones
}




func ListadoElementosCarpeta(dirPath string) ([]string, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, filepath.Join(dirPath, file.Name()))
		}
	}

	return fileNames, nil
}
