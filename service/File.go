package service

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"

	"github.com/huoxue1/filecloud/config"
)

func WriterToFiles(f io.Reader) ([]string, error) {
	buffer := bytes.NewBufferString("")
	var result []string
	getConfig := config.GetConfig()
	var n int64 = 1
	var err error
	for n > 0 && err == nil {
		n, err = io.CopyN(buffer, f, 4<<20)
		if err != nil && err.Error() == "EOF" {
			n = -1
			err = nil
		}
		s := uuid.New().String()
		file, err := os.OpenFile(fmt.Sprintf("%v/%v", getConfig.StoragePath, s), os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return result, err
		}
		_, err = file.Write(buffer.Bytes())
		if err != nil {
			return result, err
		}
		result = append(result, s)
		err = file.Close()
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBufferString("")
	}
	return result, err
}
