package proto

import (
	"bytes"
	"encoding/binary"
	"io"
)

type FrameReader struct {
	r io.Reader
	b *bytes.Buffer
}

func (fr *FrameReader) Read(p []byte) (int, error) {
	var typ FrameType
	binary.Read(fr.r, binary.BigEndian, &typ)
	println(typ.String())

	if typ == Nil {
		return 0, io.EOF
	}

	if typ == EndData {
		return 0, io.EOF
	}

	if typ == BeginData || typ == EndData {
		return fr.Read(p)
	}

	var l int64
	binary.Read(fr.r, binary.BigEndian, &l)
	if l > 0 {
		buf := make([]byte, l)
		binary.Read(fr.r, binary.BigEndian, buf)

		fr.b.Write(buf)
	}

	return fr.b.Read(p)
}

func NewReader(r io.Reader) *FrameReader {
	return &FrameReader{
		b: new(bytes.Buffer),
		r: r,
	}
}
