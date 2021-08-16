package main

import (
  "fmt"
  "container/list"
)

func main() {
  data := list.New()

  data.PushBack("Rangga")
  data.PushBack("Aziz")
  data.PushBack("Nasution")
  data.PushFront("Muhammad")

  fmt.Println(data.Front().Value)
  fmt.Println(data.Back().Value)

  for e := data.Back(); e != nil; e = e.Prev() {
    fmt.Println(e.Value)
  }
}