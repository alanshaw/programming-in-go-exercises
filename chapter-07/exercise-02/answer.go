package main

import "fmt"

func main () {
  fmt.Print("Enter a number to half: ")

  var num int
  fmt.Scanf("%d", &num)

  half, even := half(num)

  if (even) {
    fmt.Println(half, "(even)")
  } else {
    fmt.Println(half, "(odd)")
  }
}

func half (num int) (int, bool) {
  half := num / 2
  return half, half % 2 == 0
}