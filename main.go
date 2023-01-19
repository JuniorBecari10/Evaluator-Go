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
    
    if exp == "" {
      continue
    }
    
    if RunCommand(exp) {
      continue
    }
    
    tks, err := Lex(exp)
    
    if err != nil {
      continue
    }
    
    res, err := Parse(tks)
    
    if err == nil {
      fmt.Println(res)
    }
  }
}

func RunCommand(s string) bool {
  switch {
    case s == "clear" || s == "cls":
      Clear()
      return true
    
    case s == "exit":
      os.Exit(0)
    
    case s == "history" || s == "hist":
      if len(history) == 0 {
        fmt.Println("You haven't evaluated any expressions!")
        return true
      }
      
      for _, h := range history {
        fmt.Println(h)
      }
      
      return true
    
    case s == "clearvars":
      vars = make(map[string]string, 0)
      return true
    
    case s == "listvars":
      fmt.Println("List of Variables:\n")
      
      for k, v := range vars {
        fmt.Printf("%s: %s\n", k, v)
      }
      return true
  }
  
  return false
}