package service

import (
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"

	"github.com/huoxue1/filecloud/config"
)

// WriterToFiles
/**
 * @Description:
 * @param f
 * @return []string
 * @return error
 */
func WriterToFiles(f io.Reader) ([]string, error) {
	var result []string
	getConfig := config.GetConfig()
	var err error
	i := 0
	for i == 0 {
		s := uuid.New().String()
		file, err := os.OpenFile(fmt.Sprintf("%v/%v", getConfig.StoragePath, s), os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return result, err
		}
		_, err = io.CopyN(file, f, 4<<20)
		if err != nil && err.Error() == "EOF" {
			err = nil
			i = 1
		}
		result = append(result, s)
		err = file.Close()
		if err != nil {
			return nil, err
		}
	}
	return result, err
}
