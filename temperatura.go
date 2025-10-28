package main

import "fmt"

const ebulicaoFahrenheit = 212.0
const ebulicaoCelcius = 100.0
const ebulicaoKelvin = 373.15

func main() {
	var temperatura float64
	fmt.Println("Digite o número que gostaria de converter de C° para F° e K°: ")
	fmt.Scanf("%g", &temperatura)
	Conversor(temperatura)
}

func Conversor(temperatura float64) {
	conversaoFahrenheit := (temperatura * 1.8) + 32
	conversaoKelvin := (temperatura + 273.15)
	fmt.Println("A temperatura em C° é: ", temperatura)
	fmt.Println("A temperatura em F° é: ", conversaoFahrenheit)
	fmt.Println("A temperatura em K° é: ", conversaoKelvin)
}
