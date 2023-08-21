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
	result, errMsg := ReversePolish(expression)
	if errMsg != "" {
		fmt.Println("Error:", errMsg)
	}
	fmt.Printf("Result: %d\n", result)

}

func ReversePolish(expression string) (int, string) {
	tokens := strings.Fields(expression)
	stack := &stack.Stack{}

	for _, token := range tokens {
		if isOperator(token) {
			operand2 := stack.Pop()
			operand1 := stack.Pop()

			if token == "/" && operand2 == 0 {
				illegal := "WHAT YOU'VE JUST DONE IS ILLEGAL! I WILL SEE YOU IN THE COURT OF LAW"
				return 0, illegal
			}

			result := performOperation(token, operand1, operand2)
			stack.Push(result)
		} else {
			operand, err := strconv.Atoi((token))
			if err != nil {
				return 0, fmt.Sprint("Error parsing operand: %s\n", err)
			}
			stack.Push(operand)
		}
	}
	// AT THE END OF THE LOOP, THERE SHOULD BE 1 ELEMENT IN THE STACK WHICH IS THE RESULT OF THE OPERATION.
	// IF THERE MORE OR LESS THAN 1 SOMETHING WENT WRONG AND THE EXPRESSION MIGHT BE INVALID
	if stack.Size() != 1 {
		return 0, "Invalid expression"
	}

	return stack.Pop(), ""
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
