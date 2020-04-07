package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	opertion := scanner.Text()
	operator := "g"
	values := strings.Split(opertion, operator)
	value1, error1 := strconv.Atoi(values[0])
	value2, error2 := strconv.Atoi(values[1])
	if error1 != nil {
		fmt.Println(error1)
	} else {
		if error2 != nil {
			fmt.Println(error2)
		} else {
			switch operator {
			case "+":
				fmt.Println(value1 + value2)
			case "-":
				fmt.Println(value1 - value2)
			case "*":
				fmt.Println(value1 * value2)
			case "/":
				fmt.Println(value1 / value2)
			default:
				fmt.Println("Este operador no está soportado")
			}
		}
	}
}
