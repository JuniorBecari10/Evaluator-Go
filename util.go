package main

import (
  "os"
  "os/exec"
  "runtime"
)

func IsDigit(c uint8) bool {
  return c >= '0' && c <= '9'
}

func IsLetter(c uint8) bool {
  return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func Clear() {
  if (runtime.GOOS == "windows") {
    cmd := exec.Command("cmd", "/c", "cls")
    
    cmd.Stdout = os.Stdout
    cmd.Run()
  } else {
    cmd := exec.Command("clear")
    
    cmd.Stdout = os.Stdout
    cmd.Run()
  }
}

func IsKindOperator(kind int) bool {
  return kind >= 2
}

func Remove(slice []Token, i int) []Token {
    return append(slice[:i], slice[i + 1:]...)
}

func Insert(slice []Token, item Token, index int) []Token {
  slice = append(slice[:index + 1], slice[index:]...)
  slice[index] = item
  
  return slice
}