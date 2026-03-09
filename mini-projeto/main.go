package main

import (
	"fmt"
)

func main() {
	var numeros []int
	var n int

    for {
		fmt.Println("1 - Adicionar número")
		fmt.Println("2 - Listar números")
		fmt.Println("3 - Remover número por índice")
		fmt.Println("4 - Estatísticas")
		fmt.Println("5 - Divisão segura")
		fmt.Println("6 - Limpar lista")
		fmt.Println("0 - Sair")
		fmt.Print("Digite uma opção: ")

		fmt.Scanln(&n)

		if n == 1 {
			var num int
			fmt.Print("Digite um número inteiro: ")
			fmt.Scanln(&num)

			numeros = append(numeros, num)
			fmt.Println("Número adicionado com sucesso!")

		} else if n == 2 {
			if len(numeros) == 0 {
				fmt.Println("Nenhum número na lista")
			} else {
				for i := 0; i < len(numeros); i ++ {
					fmt.Println(numeros[i])
				}
			}
		} else if n == 3 {
			var index int
			fmt.Print("Digite um index para remover da lista: ")
			fmt.Scanln(&index)

			for i := range numeros {
				if i == index {
					numeros = append(numeros[:i], numeros[i+1:]...)
				}
			}
		
		} else if n == 6 {
			numeros = nil
			fmt.Println("Lista limpa com sucesso!")
		} else if n == 0 {
			break
		} else {
			fmt.Println("Opção inválida")
		}
	}
	
}