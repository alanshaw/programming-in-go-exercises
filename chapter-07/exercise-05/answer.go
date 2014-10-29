package main

import "fmt"

func main () {
  fmt.Print("Fibonacci sequence for n: ")

  var num int
  fmt.Scanf("%d", &num)

  fmt.Println(fib(num))
}

func fib (num int) int {
  if num == 0 {
    return 0
  }
  if num == 1 {
    return 1
  }
  return fib(num - 1) + fib(num - 2)
}