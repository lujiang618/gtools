package utils

import (
	"io/ioutil"
	"os"
	"path"
)

// 只删除子目录，不删除目录
func CleanDir(dirPath string) error {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range dir {
		if err := os.RemoveAll(path.Join(dirPath, file.Name())); err != nil {
			return err
		}
	}

	return nil
}
