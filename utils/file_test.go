package utils

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestIsValidPath(t *testing.T) {
	t.Log(IsValidPath("~/test.txt"))
}

// io.TeeReader分段读取: https://www.twle.cn/t/385
func TestCalcFileMD5(t *testing.T) {
	src, err := os.Open("./date_test.go")
	tmpFile, err := ioutil.TempFile("./", "temp")
	defer func() {
		cerr := tmpFile.Close()

		if err != nil {
			err = cerr
		}

		err = src.Close()
	}()

	if err != nil {
		return
	}

	defer os.Remove(tmpFile.Name())
	tee := io.TeeReader(src, tmpFile)
	t.Log(CalcFileMD5(tee))
}
