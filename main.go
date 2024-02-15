package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fileName := args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
