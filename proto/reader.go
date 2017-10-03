package proto

import (
	"bytes"
	"io"
)

type FrameReader struct {
	done bool
	r    io.Reader
	b    *bytes.Buffer
}

func (fr *FrameReader) Read(p []byte) (int, error) {
	f, err := ReadFrame(fr.r)
	if err != nil || f.Type == Begin {
		return 0, err
	}

	if f.Type == End {
		fr.done = true
		return 0, nil
	}

	pos := 0
	max := len(p)
	read := 0

	// start with reading from the buffer
	for ; pos < max; pos++ {
		b, err := fr.b.ReadByte()
		if err == io.EOF {
			break
		}

		p[pos] = b
		read++
	}

	// continue filling up p from the current data frame
	dataLen := len(f.Data)
	dataPos := 0
	for ; read <= max && dataPos < dataLen; dataPos++ {
		p[pos] = f.Data[dataPos]
		read++
		pos++
	}

	// buffer the remainder of th data frame for later
	for ; dataPos < dataLen; dataPos++ {
		fr.b.WriteByte(f.Data[dataPos])
	}

	if read == 0 && fr.done {
		return 0, io.EOF
	}

	return read, nil
}

func NewReader(r io.Reader) *FrameReader {
	return &FrameReader{
		b: new(bytes.Buffer),
		r: r,
	}
}
