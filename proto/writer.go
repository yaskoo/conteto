package proto

import (
	"encoding/binary"
	"io"
)

type FrameWriter struct {
	w io.Writer
}

func (fw *FrameWriter) Write(p []byte) (n int, err error) {
	binary.Write(fw.w, binary.BigEndian, Data)

	l := len(p)
	binary.Write(fw.w, binary.BigEndian, int64(l))
	binary.Write(fw.w, binary.BigEndian, p)
	return l, nil
}

func NewWriter(w io.Writer) *FrameWriter {
	return &FrameWriter{
		w: w,
	}
}
