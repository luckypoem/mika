package protocols

type Protocol interface {
	Write(b []byte) (n int, err error)
	Read(b []byte) (n int, err error)
	Close() error
}

func Pipe(dst, src Protocol) {
	// var buf = leakyBuf.Get()
	var buf = make([]byte, 4096)

	defer func() {
		// leakyBuf.Put(buf)
		dst.Close()
	}()

	var rerr, werr error
	var n int
	for {
		n, rerr = src.Read(buf)

		if n > 0 {
			_, werr = dst.Write(buf[:n])
		}
		if rerr != nil || werr != nil {
			return
		}
	}
}
