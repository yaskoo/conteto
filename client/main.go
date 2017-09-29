package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"

	"github.com/yaskoo/conteto/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1988")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	f, err := os.Open("./main.go")
	if err != nil {
		panic(err)
	}

	binary.Write(conn, binary.BigEndian, proto.BeginData)

	w := proto.NewWriter(conn)
	io.Copy(w, f)

	println("after copy")
	binary.Write(conn, binary.BigEndian, proto.EndData)

	var res proto.FrameType
	binary.Read(conn, binary.BigEndian, &res)
	println("ok")
}
