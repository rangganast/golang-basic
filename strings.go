package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(strings.Contains("Rangga Aziz", "Budi"))
  fmt.Println(strings.Split("Rangga Aziz", " "))
  fmt.Println(strings.ToLower("Rangga Aziz Nasution"))
  fmt.Println(strings.ToUpper("Rangga Aziz Nasution"))
  fmt.Println(strings.ToTitle("rangga aziz nasution"))
  fmt.Println(strings.Trim("         Rangga Aziz   ", " "))
  fmt.Println(strings.ReplaceAll("Rangga Rangga", "Rangga", "Eko"))
}