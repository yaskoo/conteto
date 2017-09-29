package util

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

func Id() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func MkFile(root string, repo string, id string) (*os.File, error) {
	path := fmt.Sprintf("%s/%s/%s", root, repo, time.Now().Format("2006/01/02/15/04"))

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}

	filepath := fmt.Sprintf("%s/%s", path, id)
	return os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
}
