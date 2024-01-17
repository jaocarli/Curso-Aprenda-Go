// := é uma variável chamada Gopher, usado para declarar variáveis. Deve ser usado somente dentro de um codeblock
// = é um sinal de atribuição

package main

import (
	"fmt"
)

func main() {

	x := 10
	y := "Olá, mundo!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
}
