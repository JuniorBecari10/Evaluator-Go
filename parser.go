package main

import (
  "fmt"
  "strconv"
  "math"
)

var vars map[string]int

func Parse(tks []Token) float64 {
  s := Eval(tks)
  res := PostFix(s)
  
  return res
}

// Shunting Yard Algorithm
func Eval(tks []Token) *Queue {
  stack := NewStack([]string {})
  queue := NewQueue([]string {})
  
  i := 0
  for i < len(tks) {
    // if it's a number, automatically enqueue it.
    if tks[i].kind == T_NUMBER {
      queue.Enqueue(tks[i].content)
    } else {
      // otherwise, if the length of the stack is greater than 0 (i.e. there's something in there),
      if stack.Len() > 0 {
        if tks[i].content == "(" {
          stack.Push(tks[i].content)
        } else if tks[i].content == ")" {
          pop := stack.Pop()
          count := 0
          for pop == "(" {
            if pop == ")" {
              count++
              continue
            }
            
            if pop == "(" && count > 0 {
              count--
              continue
            }
            
            queue.Enqueue(pop)
            pop = stack.Pop()
          }
        }
        
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
  
  // enqueue the remaining items in the stack
  for stack.Len() > 0 {
    queue.Enqueue(stack.Pop())
  }
  
  return queue
}

func PostFix(q *Queue) float64 {
  stack := NewStack([]string {})
  
  i := 0
  for i < q.Len() {
    if IsDigit(q.Get(i)[0]) {
      stack.Push(q.Get(i))
    } else {
      sec := stack.Pop()
      fir := stack.Pop()
      
      stack.Push(fmt.Sprint(Calc(fir, sec, q.Get(i))))
    }
    
    i++
  }
  
  v, _ := strconv.ParseFloat(stack.Pop(), 64)
  
  return v
}

func Precedence(c string) int {
  switch {
    case c == "+" || c == "-":
      return 0
    
    case c == "*" || c == "/":
      return 1
    
    case c == "^":
      return 2
  }
  
  return -1
}

func Calc(x string, y string, op string) float64 {
  a, _ := strconv.ParseFloat(x, 64)
  b, _ := strconv.ParseFloat(y, 64)
  
  switch op {
    case "+":
      return a + b
    
    case "-":
      return a - b
    
    case "*":
      return a * b
    
    case "/":
      return a / b
    
    case "^":
      return math.Pow(a, b)
  }
  
  return -1
}