package service

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestWriterToFiles(t *testing.T) {
	b, _ := os.ReadFile("lib.exe")
	reader := bytes.NewReader(b)
	files, err := WriterToFiles(reader)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
