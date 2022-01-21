package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"reflect"
	"unsafe"
)

// BytesToString 没有内存开销的转换
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytes 没有内存开销的转换
func StringToBytes(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

func CheckFile(path string) {
	_, err := os.Stat(path)
	if err != nil {
		_, err := os.Create(path)
		if err != nil {
			return
		}
		return
	}
}

func CheckDir(path string) {
	_, err := os.Stat(path)
	if err != nil {
		err := os.Mkdir(path, 0666)
		if err != nil {
			return
		}
		return
	}
}

// GetFileMd5 获取文件的md5码
func GetFileMd5(f io.Reader) string {
	md5h := md5.New()
	io.Copy(md5h, f)
	return hex.EncodeToString(md5h.Sum(nil))
}
