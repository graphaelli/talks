var (
	counts = make(chan int)
	errc   = make(chan error)
	wg     sync.WaitGroup
)
for i := 0; i < Concurrency; i++ {
	wg.Add(1) // HL
	go func() {
		defer wg.Done() // HL
		for off := range offsets {
			err := errors.New("sentinel")
			for fail := 0; err != nil && fail < 3; fail++ {
				err = getChunk(out, source, off, size, counts, rp)
			}
			if err != nil {
				errc <- fmt.Errorf("fetching chunk at offset %v: %v", off, err)
			}
		}
	}()
}
go func() {
	wg.Wait() // HL
	close(counts)
}()                                                                             
