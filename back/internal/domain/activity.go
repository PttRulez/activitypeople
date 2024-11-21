package domain

import "time"

type Activity struct {
	Calories int

	Description string

	// расстояние в километрах
	Distance int

	Date time.Time

	// общий подъём в метрах
	Elevate int

	// средний пульс
	Heartrate int
	Id        int
	Name      string

	// темп в секундах
	Pace       int
	PaceString string
	Source     ActivitySource

	// айдишник активити у поставщика (гармин, страва и т.д.)
	SourceId  int64
	SportType SportType

	// время в секундах
	TotalTime int
	UserId    int
}

type ActivityFilters struct {
	From   time.Time
	Source ActivitySource
	Until  time.Time
}

type SportType string

const (
	STOther      SportType = "Other"
	STRide       SportType = "Ride"
	STRollerSkis SportType = "RollerSkis"
	STRun        SportType = "Run"
	STStrength   SportType = "Strength"
	STXCSki      SportType = "XCSki"
)

type ActivitySource string

const (
	Strava = "strava"
)
