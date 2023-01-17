package main

import (
  "fmt"
  "strconv"
  "errors"
)

const (
  T_NUMBER = iota
  T_IDENT
  T_PLUS
  T_MINUS
  T_TIMES
  T_DIV
  T_LPAREN
  T_RPAREN
  T_EQUALS
  T_SEMICOLON
)

type Any interface {}

type Token struct {
  kind int
  content Any
  pos int
}

func Lex(s string) ([]Token, error) {
  tokens := []Token {}
  
  i := 0
  for i < len(s) {
    
    if s[i] == ' ' {
      i++
      continue
    } else if IsDigit(s[i]) {
      start := i
      
      for i < len(s) && IsDigit(s[i]) {
        i++
      }
      
      n, _ := strconv.Atoi(s[start:i])
      
      tokens = append(tokens, Token { T_NUMBER, n, start })
    } else if IsLetter(s[i]) {
      start := i
      
      for i < len(s) && IsLetter(s[i]) {
        i++
      }
      
      tokens = append(tokens, Token { T_IDENT, s[start:i], start })
    } else if s[i] == '+' {
      tokens = append(tokens, Token { T_PLUS, s[start:i], start })
    } else if s[i] == '-' {
      tokens = append(tokens, Token { T_MINUS, s[start:i], start })
    } else if s[i] == '*' {
      tokens = append(tokens, Token { T_TIMES, s[start:i], start })
    } else if s[i] == '/' {
      tokens = append(tokens, Token { T_DIV, s[start:i], start })
    } else if s[i] == '(' {
      tokens = append(tokens, Token { T_LPAREN, s[start:i], start })
    } else if s[i] == ')' {
      tokens = append(tokens, Token { T_RPAREN, s[start:i], start })
    } else if s[i] == '=' {
      tokens = append(tokens, Token { T_EQUALS, s[start:i], start })
    } else if s[i] == ';' {
      tokens = append(tokens, Token { T_SEMICOLON, s[start:i], start })
    } else {
      fmt.Printf("Unknown token: %s\n", s)
      
      // 15 times
      fmt.Printf("               ")
      
      for j := 0; j < i - 1; j++ {
        fmt.Printf(" ")
      }
      
      if i > 0 {
        fmt.Printf(" ")
      }
      
      fmt.Println("^")
      
      return nil, errors.New("unknown token")
    }
    
    i++
  }
  
  return tokens, nil
}