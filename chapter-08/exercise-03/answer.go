package main

import "fmt"

func main () {
  // Allocate enough memory for an int and return a pointer to it
  num := new(int)

  // Initially the value of the space is zero
  fmt.Println(*num)

  // Add the two numbers tongether, sorting the result in the
  // space pointed to by num
  add(1000, 138, num)

  // Log out the result
  fmt.Println(*num)
}

func add (a int, b int, res *int) {
  *res = a + b
}