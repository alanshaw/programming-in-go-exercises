package main

import "fmt"

func main () {
  var num int = 0
  inc(&num)
  fmt.Println(num)
}

func inc (num *int) {
  *num = *num + 1
}