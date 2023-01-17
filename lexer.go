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
  T_POWER
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
      i--
    } else if IsLetter(s[i]) {
      start := i
      
      for i < len(s) && IsLetter(s[i]) {
        i++
      }
      
      tokens = append(tokens, Token { T_IDENT, s[start:i], start })
      i--
    } else if s[i] == '+' {
      tokens = append(tokens, Token { T_PLUS, s[i], i })
    } else if s[i] == '-' {
      tokens = append(tokens, Token { T_MINUS, s[i], i })
    } else if s[i] == '*' {
      tokens = append(tokens, Token { T_TIMES, s[i], i })
    } else if s[i] == '/' {
      tokens = append(tokens, Token { T_DIV, s[i], i })
    } else if s[i] == '^' {
      tokens = append(tokens, Token { T_POWER, s[i], i })
    } else if s[i] == '(' {
      tokens = append(tokens, Token { T_LPAREN, s[i], i })
    } else if s[i] == ')' {
      tokens = append(tokens, Token { T_RPAREN, s[i], i })
    } else if s[i] == '=' {
      tokens = append(tokens, Token { T_EQUALS, s[i], i })
    } else if s[i] == ';' {
      tokens = append(tokens, Token { T_SEMICOLON, s[i], i })
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