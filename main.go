package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  
  fmt.Println("Expression Evaluator\n")
  
  for {
    fmt.Printf("> ")
    scanner.Scan()
    
    tks, err := Lex(scanner.Text())
    
    if err != nil {
      continue
    }
    
    fmt.Println(tks)
  }
}