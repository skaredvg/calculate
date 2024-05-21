package main

import (
	"calculator1/calc"
	"fmt"
)

func main() {
	var (
		op1, op2  float64
		operation string
	)

	fmt.Print("Операнд1: ")
	fmt.Scanln(&op1)

	fmt.Print("Операнд2: ")
	fmt.Scanln(&op2)

	fmt.Print("Операция: ")
	fmt.Scanln(&operation)

	calculator := calc.NewCalculator()
	result, err := calculator.Calculate(op1, op2, operation)

	if err != nil {
		fmt.Printf("Ошибка операции %f %s %f! %s", op1, operation, op2, err.Error())
	} else {
		fmt.Printf("Результат операции %f %s %f: %f", op1, operation, op2, result)
	}

}
