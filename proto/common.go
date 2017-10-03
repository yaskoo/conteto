package proto

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	Nil FrameType = iota
	Begin
	End
	Ok
	Error
	Meta
	Data
)

type FrameType int64

type Frame struct {
	Type   FrameType
	Length int64
	Data   []byte
}

func (f FrameType) String() string {
	var s string
	switch f {
	case Nil:
		s = "Nil"
	case Begin:
		s = "Begin"
	case End:
		s = "End"
	case Ok:
		s = "Ok"
	case Error:
		s = "Error"
	case Meta:
		s = "Meta"
	case Data:
		s = "Data"
	default:
		s = "some shit frame"
	}
	return fmt.Sprintf("%s (%d)", s, f)
}

func ReadFrame(r io.Reader) (*Frame, error) {
	var f Frame
	if err := binary.Read(r, binary.BigEndian, &f.Type); err != nil {
		return nil, err
	}

	if f.Type == Begin || f.Type == End || f.Type == Nil {
		return &f, nil
	}

	if err := binary.Read(r, binary.BigEndian, &f.Length); err != nil {
		return nil, err
	}

	data := make([]byte, int(f.Length))
	if err := binary.Read(r, binary.BigEndian, data); err != nil {
		return nil, err
	}

	f.Data = data
	return &f, nil
}
