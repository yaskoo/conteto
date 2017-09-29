package storage

import (
	"encoding/binary"
	"io"

	"github.com/spf13/viper"

	"github.com/yaskoo/conteto/proto"
	"github.com/yaskoo/conteto/util"
)

func Save(r io.Reader, w io.Writer) error {
	id := util.Id()

	file, err := util.MkFile(viper.GetString("data.location"), "my_repo", id)
	if err != nil {
		return err
	}
	defer file.Close()

	fr := proto.NewReader(r)

	_, err = io.Copy(file, fr)
	println("after copy")

	binary.Write(w, binary.BigEndian, proto.Ok)
	return err
}
