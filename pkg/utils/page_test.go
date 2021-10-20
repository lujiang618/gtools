package utils

import "testing"

func TestGetPageOffset(t *testing.T) {
	t.Log(GetPageOffset(10, 1))
	t.Log(GetPageOffset(10, 0))
	t.Log(GetPageOffset(10, 2))
	t.Log(GetPageOffset(10, 20))
}
