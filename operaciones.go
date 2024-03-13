package calculadora

import (
	"fmt"
	"time"
)

type Operacion struct {
	A                  float32
	B                  float32
	Operador           string
	CargaComputacional int
	Resultado          float32
}

// singleThread
func (o *Operacion) Operar() {

	//fmt.Println("EJECUTANDO => ST")
	time.Sleep(time.Duration(o.CargaComputacional * 100 * int(time.Millisecond)))

	if o.Operador == "+" {
		o.Resultado = o.A + o.B
	} else if o.Operador == "-" {
		o.Resultado = o.A - o.B
	} else if o.Operador == "*" {
		o.Resultado = o.A * o.B
	} else if o.Operador == "/" {
		o.Resultado = o.A / o.B
	} else if o.Operador == "%" {
		o.Resultado = o.A / o.B
	}
}

func (o Operacion) String() string {
	var cadenaSalida string
	cadenaSalida += "----------------------\n"
	cadenaSalida += fmt.Sprintf("A: %.1f \n", o.A)
	cadenaSalida += fmt.Sprintf("B: %.1f \n", o.B)
	cadenaSalida += fmt.Sprintf("Operador: %s \n", o.Operador)
	cadenaSalida += fmt.Sprintf("CargaComputacional: %d*100 ms\n", o.CargaComputacional)
	cadenaSalida += fmt.Sprintf("Resultado: %.1f \n", o.Resultado)
	cadenaSalida += "----------------------\n"
	return cadenaSalida
}
