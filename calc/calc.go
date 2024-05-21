package calc

import "fmt"

type calculator struct {
}

func NewCalculator() calculator {
	return calculator{}
}

func (c calculator) Calculate(op1, op2 float64, operation string) (result float64, err error) {
	switch operation {
	case "+":
		result, err = c.add(op1, op2)
	default:
		err = fmt.Errorf("Операция %s неподдерживается", operation)
	}
	return
}

func (c calculator) add(op1, op2 float64) (float64, error) {
	return (op1 + op2), nil
}
