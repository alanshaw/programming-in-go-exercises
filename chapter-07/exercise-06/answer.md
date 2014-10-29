`defer` allows you to schedule a function to be run after the function completes.

`panic` is a function that causes a runtime error. It causes any function to complete immediately (without executing any further statements).

`recover` is a function that can be called to recover from a panic. When called, it returns the value passed ot the original call to panic.

To recover from a panic within a function you need to defer a function that calls `recover` because `panic` causes execution of the function body to cease.