package server

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/yaskoo/conteto/storage"
)

func Start() {
	bind := viper.GetString("bind")

	l, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer l.Close()

	log.Debugf("listening on %s", bind)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Errorf("error accepting connection: %s", err.Error())
		}

		go onRequest(conn)
	}
}

func onRequest(conn net.Conn) {
	defer conn.Close()

	if err := storage.Save(conn, conn); err != nil {
		log.Errorf("unable to save: %s", err.Error())
	}
}
