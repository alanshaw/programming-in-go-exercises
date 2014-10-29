Using `makeEvenGenerator` as an example, write a `makeOddGenerator` function that generates odd numbers.

```go
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}
```