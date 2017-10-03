package testutil

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func FramedReader(src io.Reader, meta map[string]string) io.Reader {
	dest := new(bytes.Buffer)

	binary.Write(dest, binary.BigEndian, int64(1)) // Begin

	if meta != nil {
		for k, v := range meta {
			binary.Write(dest, binary.BigEndian, int64(5)) // Meta

			data := []byte(fmt.Sprintf("%s:%s", k, v))

			binary.Write(dest, binary.BigEndian, int64(len(data)))
			binary.Write(dest, binary.BigEndian, data)
		}

		if len(meta) > 0 {
			binary.Write(dest, binary.BigEndian, int64(2)) // End
		}
	}

	data := make([]byte, 256)
	for {
		n, err := src.Read(data)
		if err != nil {
			break
		}

		if n > 0 {
			binary.Write(dest, binary.BigEndian, int64(6)) // Data
			binary.Write(dest, binary.BigEndian, int64(n))

			var real_data []byte
			if n < 256 {
				real_data = make([]byte, n)
				copy(real_data, data)
			} else {
				real_data = data
			}

			binary.Write(dest, binary.BigEndian, real_data)
		}
	}

	binary.Write(dest, binary.BigEndian, int64(2))

	binary.Write(dest, binary.BigEndian, int64(2))

	return dest
}
