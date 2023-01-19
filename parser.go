package main

import (
  "fmt"
  "strconv"
  "errors"
  "math"
)

var vars map[string]string = make(map[string]string, 0)
var history []string

func Parse(tks []Token) (float64, error) {
  var_decl := false
  var_name := ""
  
  for i, t := range tks {
    // check if it's a variable declaration
    if t.kind == T_EQUALS {
      // check syntax error (i.e. the equals sign isn't in the second index inside the array and there is no identifier before it)
      if i != 1 || (i > 0 && tks[i - 1].kind != T_IDENT) {
        fmt.Println("Syntax error when declaring a variable.\n")
        fmt.Println("Examples:\nx = 10\ny = 3 - (2 * 5)")
        return -1, errors.New("var decl")
      }
      
      var_decl = true
      var_name = tks[i - 1].content
      
      // verify if the user has used a special variable name
      if var_name == "ans" {
        fmt.Printf("You cannot use '%s' as a variable name; it's a special variable.\n", var_name)
        return -1, errors.New("var special")
      }
    }
    
    // check if the user's using a variable in the expression
    // check if there's no equals after the identifier or if it's the last element in the array of tokens
    if t.kind == T_IDENT && ((i < len(tks) - 1 && tks[i + 1].kind != T_EQUALS) || (i == len(tks) - 1)) {
      val, ok := vars[t.content]
      
      // if the variable doesn't exist, show error
      if !ok {
        if t.content == "ans" {
          fmt.Println("Run an expression first to use this variable.")
        } else {
          fmt.Printf("Variable '%s' doesn't exist.\n", t.content)
        }
        return -1, errors.New("var doesn't exist")
      }
      
      // replace identifier by the variable's value
      tks[i] = Token {T_NUMBER, val, t.pos }
    }
    /*
    if i < len(tks) - 1 && t.kind == T_MINUS && tks[i + 1].kind == T_NUMBER {
      val := "-" + tks[i + 1].content
      
      tks[i + 1].content = val
      tks = Remove(tks, i)
    }*/
  }
  
  // remove variable declaration part
  if var_decl {
    tks = tks[2:]
  }
  
  fmt.Println(tks)
  
  // do evaluation
  s := Eval(tks)
  res, err := PostFix(s)
  
  if err != nil {
    return -1, err
  }
  
  // if it's a variable declaration, put the result in the variables map
  if var_decl {
    vars[var_name] = fmt.Sprint(res)
  }
  
  // update value of the special variable 'ans'
  vars["ans"] = fmt.Sprint(res)
  
  // append result to history
  history = append(history, fmt.Sprint(res))
  
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