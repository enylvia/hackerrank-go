package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	dt := make(map[string]string)
	for i := 0; i < count; i++ {
		scanner.Scan()

		dt[strings.Split(scanner.Text(), " ")[0]] = strings.Split(scanner.Text(), " ")[1]
	}
	for scanner.Scan() {
		if dt[scanner.Text()] == "" {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%s=%s\n", scanner.Text(), dt[scanner.Text()])
		}

	}
}
