package ringbuffer

import "io"

type RingBuffer struct {
	buf  []byte
	r    int
	w    int
	size int
}

var _ io.Reader = (*RingBuffer)(nil)
var _ io.Writer = (*RingBuffer)(nil)

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buf: make([]byte, size),
	}
}

func (rb *RingBuffer) Read(p []byte) (n int, err error) {
	if rb.size == 0 {
		return 0, io.EOF
	}

	toRead := len(p)
	if toRead > rb.size {
		toRead = rb.size
	}

	for i := 0; i < toRead; i++ {
		p[i] = rb.buf[rb.r]
		rb.r = (rb.r + 1) % len(rb.buf)
	}

	rb.size -= toRead
	return toRead, nil
}

func (rb *RingBuffer) Write(p []byte) (n int, err error) {
	free := len(rb.buf) - rb.size
	if free == 0 {
		return 0, io.ErrShortWrite
	}

	toWrite := len(p)
	if toWrite > free {
		toWrite = free
		err = io.ErrShortWrite
	}

	for i := 0; i < toWrite; i++ {
		rb.buf[rb.w] = p[i]
		rb.w = (rb.w + 1) % len(rb.buf)
	}

	rb.size += toWrite
	return toWrite, err
}

func (rb *RingBuffer) Len() int {
	return rb.size
}
