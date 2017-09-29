package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	_ "os"
	"testing"
)

func TestReader(t *testing.T) {
	// src, _ := os.Open("./common.go")
	src := new(bytes.Buffer)
	write_test_data(src)

	dest := new(bytes.Buffer)

	NewReader(src)

	if _, err := io.Copy(dest, src); err != nil {
		t.Error(err.Error())
	}

	fmt.Println(string(dest.Bytes()))
}

func write_test_data(w *bytes.Buffer) {
	binary.Write(w, binary.BigEndian, BeginData)
	binary.Write(w, binary.BigEndian, Data)

	data := []byte("hello world\n")
	l := int64(len(data))

	binary.Write(w, binary.BigEndian, l)
	binary.Write(w, binary.BigEndian, data)
	binary.Write(w, binary.BigEndian, EndData)
}
