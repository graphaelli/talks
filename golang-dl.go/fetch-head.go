func getSize(url string, rp RequestPreparer) (int64, error) {
	req, err := http.NewRequest("HEAD", url, nil) // HL
	if err != nil {
		return 0, fmt.Errorf("newrequest error: %v", err)
	}
	if rp != nil {
		rp(req)
	}
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return 0, fmt.Errorf("roundtrip error: %v", err)
	}
	size, err := strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parseInt error: %v", err)
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.New("ranges not supported")
	}
	return size, nil
}
