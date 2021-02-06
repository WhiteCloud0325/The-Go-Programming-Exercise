package main

import "io"

type limitReader struct {
	r        io.Reader
	n, limit int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	n, err = lr.r.Read(p[:int(lr.limit)])
	lr.n += int64(n)
	if lr.n >= lr.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r: r, limit: n}
}
