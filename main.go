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
	operator1, error1 := strconv.Atoi(values[0])
	operator2, error2 := strconv.Atoi(values[1])
	if error1 != nil {
		fmt.Println(error1)
	} else {
		if error2 != nil {
			fmt.Println(error2)
		} else {
			fmt.Println(operator1 + operator2)
		}
	}
}
