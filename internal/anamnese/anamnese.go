package anamnese

import "fmt"

// CalcularIMC calcula o IMC dado peso e altura
func CalcularIMC(peso, altura float64) float64 {
	return peso / (altura * altura)
}

// MostrarFicha mostra os dados e o IMC
func MostrarFicha(nome string, idade int, peso, altura float64) {
	fmt.Println("Olá, mundo! 🇺🇿! 👋")

	imc := CalcularIMC(peso, altura)
	fmt.Printf("Nome: %s\nIdade: %d\nPeso: %.2f kg\nAltura: %.2f m\nIMC: %.2f\n", nome, idade, peso, altura, imc)
}