package calc

type Calculator struct {
	str string
}

func NewCalculator(str string) Calculator {
	return Calculator{str}
}

func (c Calculator) Calculate() int {
	return calculate(c.str)
}

//Function that takes in an equation string like 3(2+1)-9*2^3
//and returns the result of the equation
func calculate(equ string) int {
	var stack []int
	var operatorStack []byte
	for i := 0; i < len(equ); i++ {
		if isNumber(equ[i]) {
			num := 0
			for i < len(equ) && isNumber(equ[i]) {
				num = num*10 + int(equ[i]-'0')
				i++
			}
			i--
			stack = append(stack, num)
		} else if isOperator(equ[i]) {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != '(' && getPrecedence(equ[i]) <= getPrecedence(operatorStack[len(operatorStack)-1]) {
				stack = append(stack, performOperation(operatorStack[len(operatorStack)-1], stack[len(stack)-2], stack[len(stack)-1]))
				operatorStack = operatorStack[:len(operatorStack)-1]
				stack = stack[:len(stack)-2]
			}
			operatorStack = append(operatorStack, equ[i])
		} else if isLeftBracket(equ[i]) {
			operatorStack = append(operatorStack, equ[i])
		} else if isRightBracket(equ[i]) {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != '(' {
				stack = append(stack, performOperation(operatorStack[len(operatorStack)-1], stack[len(stack)-2], stack[len(stack)-1]))
				operatorStack = operatorStack[:len(operatorStack)-1]
				stack = stack[:len(stack)-2]
			}
			if len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] == '(' {
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
		}
	}
	for len(operatorStack) > 0 {
		stack = append(stack, performOperation(operatorStack[len(operatorStack)-1], stack[len(stack)-2], stack[len(stack)-1]))
		operatorStack = operatorStack[:len(operatorStack)-1]
		stack = stack[:len(stack)-2]
	}
	return stack[0]
}

func performOperation(op byte, a int, b int) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	case '^':
		return power(a, b)
	}
	return 0
}

func power(a int, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}

func getPrecedence(c byte) int {
	switch c {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	case '^':
		return 3
	}
	return -1
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isOperator(c byte) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '^'
}

func isLeftBracket(c byte) bool {
	return c == '('
}

func isRightBracket(c byte) bool {
	return c == ')'
}

func isSpace(c byte) bool {
	return c == ' '
}
