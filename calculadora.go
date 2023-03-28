package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {
}

func (Calc) Operate(entrada string, operador string) int {

	valores := strings.Split(entrada, operador)
	valor1 := Parsear(valores[0])
	valor2 := Parsear(valores[1])

	switch operador {
	case "+":
		fmt.Println(valor1 + valor2)
		return valor1 + valor2
	case "-":
		fmt.Println(valor1 - valor2)
		return valor1 - valor2
	case "*":
		fmt.Println(valor1 * valor2)
		return valor1 * valor2
	case "/":
		fmt.Println(valor1 / valor2)
		return valor1 / valor2
	default:
		fmt.Println("Invalid")
		return 0
	}
}

func Parsear(valores string) int {
	valor, err := strconv.Atoi(valores)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return valor
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
