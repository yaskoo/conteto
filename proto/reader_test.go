package proto

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"testing"

	"github.com/yaskoo/conteto/testutil"
)

func TestReader(t *testing.T) {
	src, err := os.Open("../testutil/data/hello_world")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer src.Close()

	r := NewReader(testutil.FramedReader(src, nil))

	dest := new(bytes.Buffer)
	if _, err := io.Copy(dest, r); err != nil {
		t.Error(err.Error())
	}

	expected := "hello world"
	actual := string(dest.Bytes())
	if expected != actual {
		t.Errorf("Expected: [%s], got: [%s]", expected, actual)
	}
}
