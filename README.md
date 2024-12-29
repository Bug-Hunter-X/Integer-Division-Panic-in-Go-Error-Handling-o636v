# Integer Division Panic in Go Error Handling

This repository demonstrates a subtle error in Go's error handling. The `myFunc` function attempts to handle the case where the input `x` is 0, returning an error if this occurs. However, it still calculates `10 / x` before returning the error, leading to a panic if `x` is 0. The solution shows how to avoid this error by only calculating `10 / x` if `x` is not 0.