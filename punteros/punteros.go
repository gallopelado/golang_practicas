package main

import "fmt"

func main() {
	x := 25
	fmt.Println("desde el main, su address", &x)
	fmt.Println("desde el main, su valor", x)
	cambiarValor(&x)
	fmt.Println("desde el main, su valor CAMBIADO", x)
}

func cambiarValor(a *int) {
	// misma referencia al valor
	fmt.Println("desde la funcion, su address", a)
	fmt.Println("desde la funcion, su valor", *a)
	*a = 66
	//a = 16
	fmt.Println("desde la funcion, su address", a)
	fmt.Println("desde la funcion, su valor CAMBIADO", *a)
}

/*
En Go hay dos tipos de parámetros que podemos usar en nuestras funciones o métodos:

Pass by value (por valor) : Aquí la función hara una copia de la variable que se esta pasando y modificara la copia.
Pass by pointer/reference (por puntero) : Aquí la función obtendrá un puntero que apunta a la misma dirección donde esta el valor
de la variable o struct
Usos comunes:
Pass by value es usado cuando necesitamos utilizar los valores de las variables o propiedades de un struct SIN CAMBIAR EL VALOR DE ELLAS.
Ahora, si quieres una función que SI CAMBIE EL VALOR de una propiedad, por ejemplo la edad. La funcion requiere que usemos PASS BY POINTER
*/
