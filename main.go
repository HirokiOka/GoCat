package main

import (
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

}
