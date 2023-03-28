package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (calc) operate(entrada string, operador string) int {

	valores := strings.Split(entrada, operador)
	valor1 := parsear(valores[0])
	valor2 := parsear(valores[1])

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

func parsear(valores string) int {
	valor, err := strconv.Atoi(valores)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return valor
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
