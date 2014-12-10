offsets := make(chan int64)
go func() {
	for i := int64(0); i < size; i += int64(Blocksize) {
		offsets <- i
	}
	close(offsets)
}()
