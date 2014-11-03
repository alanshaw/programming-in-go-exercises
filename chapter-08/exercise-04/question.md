What is the value of x after running this program:

```go
func square(x *float64) {
    *x = *x * *x
}
func main() {
    x := 1.5
    square(&x)
}
```