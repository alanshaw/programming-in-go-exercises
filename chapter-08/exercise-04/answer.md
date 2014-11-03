`&x` returns the memory location of `x`, since `x` is a `float`, `&x` returns a pointer to an `float`: a `*float`.

The pointer is passed to the `square` function and is dereferenced three times:

```go
*x = *x * *x
``` 

This reads as: multiply the value at memory location `x` by the value at memory location `x` and store the result in memory location `x`.

In our case, the value is 1.5, so after the call to `square`, `x` in main will have the value 2.25.