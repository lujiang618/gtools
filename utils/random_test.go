package utils

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestRandomInt(t *testing.T) {
	t.Log(RandomInt(100, 1000))
	t.Log(RandomInt(100000, 999999))
	t.Log(RandomInt(1000, 9999))
	t.Log(RandomInt(1000, 1000))
	t.Log(RandomInt(10000, 1000))
	t.Log(RandomInt(100, 1000))
}

func TestRandomString(t *testing.T) {
	t.Log(RandomString(0))
	t.Log(RandomString(4))
	t.Log(RandomString(6))
}

func Test_generator(t *testing.T) {
	tests := []struct {
		name string
		want *rand.Rand
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generator() = %v, want %v", got, tt.want)
			}
		})
	}
}
