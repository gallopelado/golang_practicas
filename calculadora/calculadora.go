package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()
	// fmt.Println(operacion)
	resultado := 0
	operador := ""
	var valores []string
	switch {
	case strings.Contains(operacion, "+"):
		valores = strings.Split(operacion, "+")
		operador = "+"
		operador1, operador2, estado := parseValores(valores[0], valores[1])
		if estado {
			resultado = operador1 + operador2
		}
	case strings.Contains(operacion, "-"):
		valores = strings.Split(operacion, "-")
		operador = "-"
		operador1, operador2, estado := parseValores(valores[0], valores[1])
		if estado {
			resultado = operador1 - operador2
		}
	case strings.Contains(operacion, "*"):
		valores = strings.Split(operacion, "*")
		operador = "*"
		operador1, operador2, estado := parseValores(valores[0], valores[1])
		if estado {
			resultado = operador1 * operador2
		}
	default:
		fmt.Println("Ningún caso aplicado")
	}
	fmt.Printf("La operación es %v, el resultado es %d\n", operador, resultado)
	//valores = strings.Split(operacion, "+") || strings.Split(operacion, "-")
	// fmt.Printf("%T\n", valores)
	// conversiones
	// operador1, err1 := strconv.Atoi(valores[0])
	// if err1 != nil {
	// 	fmt.Println("Error en operador1", err1)
	// 	return
	// }
	// operador2, err2 := strconv.Atoi(valores[1])
	// if err2 != nil {
	// 	fmt.Println("Error en operador2", err2)
	// 	return
	// }
	//operador := "+"
	// resultado := 0
	// switch {
	// case strings.Contains(operacion, "+"):
	// 	operador = "+"
	// 	resultado = operador1 + operador2
	// case strings.Contains(operacion, "-"):
	// 	operador = "-"
	// 	resultado = operador1 - operador2
	// case strings.Contains(operacion, "*"):
	// 	operador = "*"
	// 	resultado = operador1 * operador2
	// case strings.Contains(operacion, "/"):
	// 	if operador2 == 0 {
	// 		fmt.Println("Error, división sobre cero")
	// 		return
	// 	}
	// 	operador = "/"
	// 	resultado = operador1 / operador2
	// }
	// fmt.Printf("La operación es %v, el resultado es %d", operador, resultado)
	//fmt.Println(operador1 + operador2)
}

// parseValores formats two string parameters to string :)
func parseValores(num1, num2 string) (int, int, bool) {
	isValid := true
	operador1, err1 := strconv.Atoi(num1)
	if err1 != nil {
		fmt.Println("Error en operador1", err1)
		isValid = false
	}
	operador2, err2 := strconv.Atoi(num2)
	if err2 != nil {
		fmt.Println("Error en operador2", err2)
		isValid = false
	}
	return operador1, operador2, isValid
}
