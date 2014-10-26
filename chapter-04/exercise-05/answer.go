package main

import "fmt"

func main () {
  fmt.Print("Enter a temperature in farenheit: ")
  var farenheit float64
  fmt.Scanf("%f", &farenheit)

  celsius := (farenheit - 32) * 5/9

  fmt.Println(celsius, "celsius")
}