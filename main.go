package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	bFlag := flag.Bool("b", false, "Number non-blank output lines")
	nFlag := flag.Bool("n", false, "Number all output lines")
  sFlag := flag.Bool("s", false, "Squeeze multiple adjacent empty lines")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("No file to read")
		return
	}

	fileNames := make([]string, 0)
	for _, arg := range args {
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
		lineNum := 1
		indentString := "    "
		for scanner.Scan() {
			if *bFlag {
				if scanner.Text() == "\n" || scanner.Text() == "" {
					fmt.Println(scanner.Text())
				} else {
					fmt.Println(indentString, lineNum, "", scanner.Text())
					lineNum++
				}
			} else if *nFlag {
				fmt.Println(indentString, lineNum, "", scanner.Text())
				lineNum++
      } else if *sFlag {
				if scanner.Text() == "" {
          fmt.Println(scanner.Text())
          for scanner.Scan() {
            if scanner.Text() != "" {
              fmt.Println(scanner.Text())
              break
            }
          }
        } else {
          fmt.Println(scanner.Text())
        }
			} else {
				fmt.Println(scanner.Text())
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
