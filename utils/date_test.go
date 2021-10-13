package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestGetCurrentDateTime(t *testing.T) {
	t.Log(GetCurrentDateTime())
}

func TestGetMonthDay(t *testing.T) {
	t.Log(GetMonthDay())
}

func TestGetWeekday(t *testing.T) {
	t.Log(GetWeekday())
}

func TestGetDateOfWeek(t *testing.T) {
	t.Log(GetDateOfWeek(time.Now(), 3))
}

func TestGenerateLocalZone(t *testing.T) {
	tests := []struct {
		name string
		want *time.Location
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateLocalZone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateLocalZone() = %v, want %v", got, tt.want)
			}
		})
	}
}
