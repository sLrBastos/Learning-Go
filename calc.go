package main

import (
	"bufio"
	"deez/stack"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the reverse Polish expression: ")
	expression, _ := reader.ReadString('\n')
	result := ReversePolish(expression)
	fmt.Printf("Result: %d\n", result)

}

func ReversePolish(expression string) int {
	tokens := strings.Fields(expression)
	stack := &stack.Stack{}

	for _, token := range tokens {
		if isOperator(token) {
			operand2 := stack.Pop()
			operand1 := stack.Pop()
			result := performOperation(token, operand1, operand2)
			stack.Push(result)
		} else {
			operand, err := strconv.Atoi((token))
			if err != nil {
				fmt.Printf("Error parsing operand: %s\n", err)
				return 0
			}
			stack.Push(operand)
		}
	}

	if stack.Size() != 1 {
		fmt.Println("Invalid expression")
		return 0
	}

	return stack.Pop()
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func performOperation(operator string, operand1, operand2 int) int {

	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("That's illegal")
			return 0
		}
		return operand1 / operand2
	default:
		fmt.Printf("Unknown operator: %s\n", operator)
		return 0
	}
}
