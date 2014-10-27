The `[low : high]` notation allows you to create a slice from an existing array. `low` is the index to start at and `high` is the index to end (non inclusive).

Note that both `low` and `high` are optional. Omitting `low` is the same as specifying `0` and omitting `high` is the same as specifying the length of the array.

```go
x := [6]string{"a","b","c","d","e","f"}
x[2:5]
```

Therefore, the above code yields a new slice containing the elements `"c"`, `"d"` and `"e"`.