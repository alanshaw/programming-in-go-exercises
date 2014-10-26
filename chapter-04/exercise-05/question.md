Using the example program as a starting point, write a program that converts from Farenheit into Celsius. `(C = (F - 32) * 5/9)`

Example program:

```go
package main

import "fmt"

func main () {
  fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}
```