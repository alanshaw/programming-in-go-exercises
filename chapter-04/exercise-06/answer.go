package main

import "fmt"

func main () {
  fmt.Print("Enter a distance in feet: ")
  var feet float64
  fmt.Scanf("%f", &feet)

  meters := 0.3048 * feet

  fmt.Println(meters, "meters")
}