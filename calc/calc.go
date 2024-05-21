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
	case "-":
		result, err = c.sub(op1, op2)
	case "*":
		result, err = c.mul(op1, op2)
	case "/":
		result, err = c.div(op1, op2)
	default:
		err = fmt.Errorf("Операция %s неподдерживается", operation)
	}
	return
}

func (c calculator) add(op1, op2 float64) (float64, error) {
	return (op1 + op2), nil
}

func (c calculator) sub(op1, op2 float64) (float64, error) {
	return (op1 - op2), nil
}

func (c calculator) mul(op1, op2 float64) (float64, error) {
	return (op1 * op2), nil
}

func (c calculator) div(op1, op2 float64) (float64, error) {
	if op2 == 0.0 {
		return op2, fmt.Errorf("Делитель не может быть равен 0")
	}
	return (op1 / op2), nil
}
