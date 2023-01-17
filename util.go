package main

func IsDigit(c uint8) bool {
  return c >= '0' && c <= '9'
}

func IsLetter(c uint8) bool {
  return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}