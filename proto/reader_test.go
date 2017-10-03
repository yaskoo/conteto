package proto

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	src := new(bytes.Buffer)
	write_test_data(src)

	dest := new(bytes.Buffer)

	r := NewReader(src)

	if _, err := io.Copy(dest, r); err != nil {
		t.Error(err.Error())
	}

	expected := "hello world"
	actual := string(dest.Bytes())
	if expected != actual {
		t.Errorf("Expected: %s, got: %s", expected, actual)
	}
}

func write_test_data(w *bytes.Buffer) {
	binary.Write(w, binary.BigEndian, Begin)
	binary.Write(w, binary.BigEndian, Data)

	data := []byte("hello world")
	l := int64(len(data))

	binary.Write(w, binary.BigEndian, l)
	binary.Write(w, binary.BigEndian, data)
	binary.Write(w, binary.BigEndian, End)
}
