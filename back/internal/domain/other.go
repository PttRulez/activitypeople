package domain

import "time"

type CalendarTime struct {
	time.Time
}

type TimeFilters struct {
	From  time.Time
	Until time.Time
}
