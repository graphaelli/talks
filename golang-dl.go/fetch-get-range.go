func getChunk(w io.WriterAt, url string, off, size int64, counts chan int, rp RequestPreparer) error {
	req, err := http.NewRequest("GET", url, nil) // HL
	if err != nil {
		return fmt.Errorf("newrequest error: %v", err)
	}
	if rp != nil {
		rp(req)
	}
	end := off + int64(Blocksize) - 1
	if end > size-1 {
		end = size - 1
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", off, end))

	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {                                         OMIT
		return fmt.Errorf("roundtrip error: %v", err)       OMIT
	}                                                       OMIT
	if res.StatusCode != http.StatusPartialContent {        OMIT
		return fmt.Errorf("bad status: %v", res.Status)     OMIT
	}                                                       OMIT
	wr := fmt.Sprintf("bytes %v-%v/%v", off, end, size)     OMIT
	if cr := res.Header.Get("Content-Range"); cr != wr {    OMIT
		res.Body.Close()                                    OMIT
		return fmt.Errorf("bad content-range: %v", cr)      OMIT
	}                                                       OMIT

	_, err = io.Copy(&sectionWriter{w, off}, logReader{res.Body, counts})  // HL
	res.Body.Close()
	if err != nil {
		return fmt.Errorf("copy error: %v", err)
	}

	return nil
}                   
