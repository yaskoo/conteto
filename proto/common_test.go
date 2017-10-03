package proto

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestReadNoDataFrames(t *testing.T) {
	empty_data := []byte{}

	writeAndAssertReadFrame(t, Nil, empty_data)
	writeAndAssertReadFrame(t, Begin, empty_data)
	writeAndAssertReadFrame(t, End, empty_data)

}

func TestReadDataFrames(t *testing.T) {
	writeAndAssertReadFrame(t, Meta, []byte("hello world"))
	writeAndAssertReadFrame(t, Data, []byte("hello world"))
	writeAndAssertReadFrame(t, Ok, []byte("hello world"))
	writeAndAssertReadFrame(t, Error, []byte("hello world"))
}

func writeAndAssertReadFrame(t *testing.T, typ FrameType, data []byte) {
	rw := new(bytes.Buffer)

	binary.Write(rw, binary.BigEndian, typ)

	data_len := int64(len(data))
	if data_len > 0 {
		binary.Write(rw, binary.BigEndian, data_len)
		binary.Write(rw, binary.BigEndian, data)
	}

	f, err := ReadFrame(rw)
	if err != nil {
		t.Error(err)
	}

	if f.Type != typ {
		t.Errorf("Expected: %s, got: %s", typ, f.Type)
	}

	if data_len == 0 {
		return
	}

	if f.Length != data_len {
		t.Errorf("Expected: %d, got: %d", data_len, f.Length)
	}

	sum := sha256.Sum256(data)
	expected_sum := fmt.Sprintf("%x", sum)

	sum = sha256.Sum256(f.Data)
	actual_sum := fmt.Sprintf("%x", sum)

	if expected_sum != actual_sum {
		t.Errorf("Expected: %s, got: %s", expected_sum, actual_sum)
	}
}
