package main

import "fmt"

const ebulicaoK = 373.0

func main() {

	tempK := ebulicaoK
	tempC := tempK - 273.0

	fmt.Println("A temperatura em ebulição da água em °K = ", tempK)
	fmt.Println("A temperatura em ebulição da água em °C = ", tempC)

}
