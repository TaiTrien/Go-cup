package util

import (
	"fmt"
	"time"
)

func DateFormat(t time.Time) string {
	return fmt.Sprintf("%s, %d-%02d-%02d %02d:%02d",
		t.Weekday(), t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute())
}
