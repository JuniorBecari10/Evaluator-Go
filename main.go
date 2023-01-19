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
      ans := vars["ans"]
      vars = make(map[string]string, 0)
      
      vars["ans"] = ans
      return true
    
    case s == "listvars":
      for k, v := range vars {
        fmt.Printf("%s: %s\n", k, v)
      }
      return true
    
    case s == "help":
      fmt.Println("Commands:")
      fmt.Println("clear/cls - Clear the screen;")
      fmt.Println("history/hist - List the history of results;")
      fmt.Println("clearvars - Clear all variables;")
      fmt.Println("listvars - List all variables;")
      fmt.Println("help - Display this help message.")
      
      return true
  }
  
  return false
}