package main

import "fmt"

var num1 float64
var num2 float64
var operacao string

func soma(a, b float64) float64 {
	return a + b
}

func subtracao(a, b float64) float64 {
	return a - b
}

func multiplicacao(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	return a / b
}

func main() {

	fmt.Println("=== CALCULADORA BÁSICA ===")
	fmt.Println("Operaçãoes: +, - , * e /. Digite 'sair' para encerrar.")

	for {
		fmt.Print("\nOperação: ")
		fmt.Scan(&operacao)

		if operacao == "sair" {
			fmt.Println("Encerrando...até mais!")
			break
		}

		fmt.Print("Primeiro número: ")
		fmt.Scan(&num1)
		fmt.Print("Segundo número: ")
		fmt.Scan(&num2)

		if operacao == "+" {
			fmt.Println("Resultado: \n", soma(num1, num2))
		} else if operacao == "-" {
			fmt.Println("Resultado: \n", subtracao(num1, num2))
		} else if operacao == "*" {
			fmt.Println("Resultado: \n", multiplicacao(num1, num2))
		} else if operacao == "/" {
			if num2 == 0 {
				fmt.Println("ERRO: Não é possível dividir nenhum numero por 0!")
			} else {
				fmt.Println("Resultado: \n", dividir(num1, num2))
			}
		} else {
			fmt.Println("Operacao Invalida.")
		}

	}
}
