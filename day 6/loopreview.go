package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < count; i++ {
		scanner.Scan()
		printIt(scanner.Text())
	}
}
func printIt(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Print(string(str[i]))
		i++
	}
	fmt.Print(" ")
	for i := 1; i < len(str); i++ {
		fmt.Print(string(str[i]))
		i++
	}

	fmt.Println("")
}
