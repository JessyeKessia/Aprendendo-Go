package main

import (
	"fmt"
	"errors"
)

func estatisticas(numeros []int) (int, int, float64, error) {
	if len(numeros) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}

	min := numeros[0]
	max := numeros[0]
	soma := 0

	for _, v := range numeros {
		soma += v
	}

	for _, v := range numeros {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	media := float64(soma) / float64(len(numeros))

	return min, max, media, nil
}

func divisao(a int, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Não pode dividir por zero")
	}

	r := float64(a) / float64(b)

	return r, nil
}



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
		
		} else if n == 4 {
			min, max, media, erro := estatisticas(numeros)

			if erro != nil {
				fmt.Println(erro)
			} else {
				fmt.Println("Mínimo:", min)
				fmt.Println("Máximo:", max)
				fmt.Printf("Média: %.2f\n", media)
			}

		} else if n == 5 {
			var n1 int
			var n2 int

			fmt.Println("Digite um número: ")
			fmt.Scanln(&n1)
			fmt.Println("Digite um segundo número: ")
			fmt.Scanln(&n2)

			r, erro := divisao(n1, n1)

			if erro != nil {
				fmt.Println(erro)
			} else {
				fmt.Printf("Divisão: %.2f\n", r)
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