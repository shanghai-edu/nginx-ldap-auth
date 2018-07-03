package utils

import (
	"testing"
)

func Test_timeCheck(t *testing.T) {
	timeRanges := []string{"11:02-17:11", "05:00"}
	b, err := TimeCheck(timeRanges)
	if err != nil {
		t.Error(err)
	}
	t.Log(b)
}
