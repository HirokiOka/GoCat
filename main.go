package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("No file to read")
		return
	}
	fileNames := make([]string, 0)

	for i, arg := range args {
		if i == 0 {
			continue
		}
		fileNames = append(fileNames, arg)

	}

	for _, fileName := range fileNames {
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
}
