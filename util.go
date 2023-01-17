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