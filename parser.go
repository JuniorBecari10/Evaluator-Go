package main

import (
  "fmt"
  "strconv"
  "errors"
  "math"
)

var vars map[string]int

func Parse(tks []Token) (float64, error) {
  s := Eval(tks)
  res, err := PostFix(s)
  
  if err != nil {
    return -1, err
  }
  
  return res, nil
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
        // check for parenthesis; if it's a opening one, push it to the stack
        if tks[i].content == "(" {
          stack.Push(tks[i].content)
          // if it's a closing one
        } else if tks[i].content == ")" {
          // pop every item of the stack until find the matching opening parenthesis
          pop := stack.Pop()
          count := 0
          // loop to find
          for pop != "(" {
            if pop == ")" {
              count++
              pop = stack.Pop()
              continue
            }
            
            if pop == "(" && count > 0 {
              count--
              pop = stack.Pop()
              continue
            }
            
            // enqueue the popped item
            queue.Enqueue(pop)
            pop = stack.Pop()
          }
          // if the precedence of the current item is greater than the one on the top of the stack, push it
        } else if Precedence(tks[i].content) > Precedence(stack.Peek()) {
          stack.Push(tks[i].content)
        } else {
          // otherwise enqueue
          queue.Enqueue(tks[i].content)
        }
        // otherwise, push
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

func PostFix(q *Queue) (float64, error) {
  stack := NewStack([]string {})
  
  i := 0
  for i < q.Len() {
    if IsDigit(q.Get(i)[0]) {
      stack.Push(q.Get(i))
    } else {
      sec := stack.Pop()
      fir := stack.Pop()
      
      // check division by zero
      if sec == "0" && q.Get(i) == "/" {
        fmt.Println("Division by zero.")
        return -1, errors.New("zero div")
      }
      
      stack.Push(fmt.Sprint(Calc(fir, sec, q.Get(i))))
    }
    
    i++
  }
  
  v, _ := strconv.ParseFloat(stack.Pop(), 64)
  
  return v, nil
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