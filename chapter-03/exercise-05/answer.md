```go
// && is true only when both operands are true
(true && false) || (false && true) || !(false && false)
// Eval the &&'s
false || false || !(false)
// ...and the ! (not)
false || false || true
// || is true when one of the operands is true
// Eval the first ||
false || true
// ... and the second
true
```