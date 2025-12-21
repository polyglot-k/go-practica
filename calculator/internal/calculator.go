package internal

import "fmt"

// Calculate는 두 개의 피연산자와 한 개의 연산자를 받아 계산 결과를 반환합니다.
func Calculate(a, b Operand, op Operator) (Operand, error) {
	switch op {
	case ADD:
		return a + b, nil
	case SUB:
		return a - b, nil
	case MUL:
		return a * b, nil
	case DIV:
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator")
	}
}