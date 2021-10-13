package utils

import (
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
