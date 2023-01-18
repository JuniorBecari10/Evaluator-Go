package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  Clear()
  
  fmt.Println("Expression Evaluator\n")
  
  for {
    fmt.Printf("> ")
    scanner.Scan()
    
    exp := scanner.Text()
    
    if RunCommand(exp) {
      continue
    }
    
    tks, err := Lex(exp)
    
    if err != nil {
      continue
    }
    
    res := Parse(tks)
    
    fmt.Println(res)
  }
}

func RunCommand(s string) bool {
  switch {
    case s == "clear" || s == "cls":
      Clear()
      return true
    
    case s == "exit":
      os.Exit(0)
  }
  
  return false
}