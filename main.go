package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calculate struct{}

func (calculate) operate(entry string, operator string) int {
	value1, value2 := splitEntry(entry, operator)
	switch operator {
	case "+":
		return value1 + value2
	case "-":
		return value1 - value2
	case "*":
		return value1 * value2
	case "/":
		return value1 / value2
	default:
		panic("Esta operación no está soportada")
	}
}

func parsing(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intValue
}

func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func splitEntry(entry string, operator string) (int, int) {
	values := strings.Split(entry, operator)
	value1 := parsing(values[0])
	value2 := parsing(values[1])
	return value1, value2
}

func main() {
	entry := readEntry()
	operator := readEntry()
	calculateInstance := calculate{}
	fmt.Println(calculateInstance.operate(entry, operator))
}
