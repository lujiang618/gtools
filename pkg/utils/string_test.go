package utils

import (
	"testing"
)

func TestStringSplit(t *testing.T) {
	var list = []string{"123", "456", "789"}
	t.Log(StringSplit(list...))
}

func TestMd5(t *testing.T) {
	str := "111111"
	t.Log(Md5(str))
}

func TestCapitalize(t *testing.T) {
	t.Log(Capitalize("asd"))
	t.Log(Capitalize("Asd"))
	t.Log(Capitalize("123"))
	t.Log(Capitalize("长江长江"))
}
