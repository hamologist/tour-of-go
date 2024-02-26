package main

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	n := len(b)

	for i := 0; i < n; i++ {
		b[i] = 'A'
	}

	return n, nil
}
