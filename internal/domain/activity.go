package domain

type Activity struct {
	// расстояние в метрах
	Distance float64

	// время в секундах
	TotalTime int

	Id        int
	Name      string
	SportType SportType
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
