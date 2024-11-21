package domain

import "time"

type DiaryDay struct {
	Activity

	Day   int
	Month time.Month
	Year  int
}
