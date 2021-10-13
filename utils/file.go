package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

//判断文件或文件夹是否存在 error为nil 则存在
func IsValidPath(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func CalcFileMD5(src io.Reader) string {
	h := md5.New()
	if n, err := io.Copy(h, src); err != nil || n == 0 {
		log.Println(err)
		return ""
	}

	return hex.EncodeToString(h.Sum(nil))
}
