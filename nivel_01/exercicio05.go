// Em package level scope, utilizando a palavra chave var, crie uma variavel com o identificador y.
// O tipo desta variavel deve ser o tipo subjacente do tipo que voce criou.

// Na funcao main.
// Demonstre o valor da varivel x
// Demonstre o tipo da variavel x
// atribua 42 a variavel x utilizando operador =
// Demonstre o valor da variavel x

// Utilize conversao para transformar o tipo do valor da variavel x em seu tipo subjacente e utilizando o operador =, atribua o valor x a y
// demosnttre o valor de y
// demontre o tipo de y

package main

import "fmt"

type pastel int

var x pastel
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
