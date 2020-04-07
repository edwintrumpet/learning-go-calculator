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
	values := strings.Split(opertion, "+")
	operator1, _ := strconv.Atoi(values[0])
	operator2, _ := strconv.Atoi(values[1])
	fmt.Println(operator1 + operator2)
}