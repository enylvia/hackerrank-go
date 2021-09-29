package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	//SOLVE
	// Declare second integer, double, and String variables.
	var (
		ii int
		dd float64
		ss string
	)
	// Read and save an integer, double, and String to your variables.
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			ii, _ = strconv.Atoi(scanner.Text())
		}
		if i == 1 {
			dd, _ = strconv.ParseFloat(scanner.Text(), 32)
		}
		if i == 2 {
			ss = scanner.Text()
		}
	}
	// Print the sum of both integer variables on a new line.
	fmt.Println(int(i) + ii)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+dd)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + ss)
}
