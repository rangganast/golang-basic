package main

import (
  "reflect"
  "fmt"
)

type Sample struct {
  Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool{
  t := reflect.TypeOf(data)
  for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    if field.Tag.Get("required") == "true" {
      value := reflect.ValueOf(data).Field(i).Interface()
      if value == "" {
        return false
      }
    }
  }
  return true
}

type ContohLagi struct {
  Name string
  Description string
}

func main() {
  sample := Sample{"Rangga"}

  var sampleType reflect.Type = reflect.TypeOf(sample)

  fmt.Println(sampleType.NumField())
  fmt.Println(sampleType.Field(0).Name)
  fmt.Println(sampleType.Field(0).Tag.Get("required"))
  fmt.Println(sampleType.Field(0).Tag.Get("max"))
  fmt.Println(sampleType.Field(0).Tag.Get("min"))

  sample.Name = ""
  fmt.Println(IsValid(sample))

  contoh := ContohLagi{"", ""}
  fmt.Println(IsValid(contoh))
}