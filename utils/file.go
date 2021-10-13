package utils

import "os"

//判断文件或文件夹是否存在 error为nil 则存在
//判断文件或文件夹是否存在 error为nil 则存在
func IsValidPath(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}
