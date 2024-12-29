func myFunc(x int) (int, error) {
    if x == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    result := 10 / x
    return result, nil
}