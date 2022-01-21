package model

import (
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestInsert(t *testing.T) {
	info := FileInfo{
		ID:         "",
		FileName:   "123",
		IsDir:      0,
		Content:    []string{"123", "234"},
		UploadTime: 123132131213,
		ParentID:   "12222",
	}
	err := InsertFile(info)
	if err != nil {
		log.Errorln(err.Error())
	}
}

func TestQueryByID(t *testing.T) {
	f := new(FileInfo)
	err := QueryByID("e7e2e1f0-5836-43ab-bdc1-d48e3e529acd", f)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	fmt.Println(f.FileName, f.Content, f.UploadTime)
}

func TestDeleteByID(t *testing.T) {
	err := DeleteByID("e7e2e1f0-5836-43ab-bdc1-d48e3e529acd")
	if err != nil {
		log.Errorln(err.Error())
		return
	}
}

func TestQueryByParent(t *testing.T) {
	fileInfos, err := QueryByParent("12222")
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	for _, f := range fileInfos {
		fmt.Println(f.FileName, f.Content, f.UploadTime)

	}
}

func TestQueryByName(t *testing.T) {
	fileInfos, err := QueryByName("4")
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	for _, f := range fileInfos {
		fmt.Println(f.FileName, f.Content, f.UploadTime)

	}
}
