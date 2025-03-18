package main

import "fmt"

//Declaração da variável CONST do ponto de ebulição da água em F
const ebulicaoF = 212.0

func main() {

	tempF := ebulicaoF            		//Var ou ":="
	tempC := (tempF - 32) * 5 / 9 		//Var ou ":="

	fmt.Println("A temperatura em ebulição da água em °F = ", tempF)
	fmt.Println("A temperatura em ebulição da água em °C = ", tempC)

}
