package main

import "fmt"

var vars map[string]int

func Parse(tks []Token) int {
  s := Eval(tks)
  fmt.Println(s.value)
  
  return 0
}

// Shunting Yard Algorithm
func Eval(tks []Token) *Queue {
  stack := NewStack([]string {})
  queue := NewQueue([]string {})
  
  i := 0
  for i < len(tks) {
    if tks[i].kind == T_NUMBER {
      queue.Enqueue(tks[i].content)
    } else {
      if stack.Len() > 0 {
        if Precedence(tks[i].content) > Precedence(stack.Peek()) {
          stack.Push(tks[i].content)
        } else {
          queue.Enqueue(tks[i].content)
        }
      } else {
         stack.Push(tks[i].content)
       }
    }
    
    i++
  }
  
  for stack.Len() > 0 {
    queue.Enqueue(stack.Pop())
  }
  
  return queue
}

func Precedence(c string) int {
  switch {
    case c == "+" || c == "-":
      return 0
    
    case c == "*" || c == "/":
      return 1
    
    case c == "^":
      return 2
    
    case c == "(" || c == ")":
      return 3
  }
  
  return -1
}