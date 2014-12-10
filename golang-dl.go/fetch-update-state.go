state := make(chan State, 1)
go func() {
    defer close(state)

    var count int64
    for {
        select {
        case n, ok := <-counts:
            count += int64(n)
            if !ok {
                err := out.Close()
                state <- State{count, err}
                return
            }
            state <- State{count, nil}
        case err := <-errc:
            out.Close()
            state <- State{count, err}
            return
        }
    }
}()

return size, state, nil
