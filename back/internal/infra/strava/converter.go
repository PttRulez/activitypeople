package strava

import (
	"fmt"
	"math"

	"github.com/pttrulez/activitypeople/internal/domain"
)

var kilometer float64 = 1000

func FromStravaToActivity(activity AthleteActivityResponse) domain.Activity {
	pace, paceString := 0, ""
	sportType := FromStravaSportType(activity.SportType)
	if sportType != domain.STStrength && activity.Distance > 0 {
		pace, paceString = calculatePace(activity.MovingTime, activity.Distance)
	}

	return domain.Activity{
		Calories:   int(activity.Calories),
		Distance:   int(activity.Distance),
		Date:       activity.StartDateLocal,
		Elevate:    int(activity.TotalElevationGain),
		Heartrate:  int(activity.AverageHeartrate),
		Name:       activity.Name,
		Pace:       pace,
		PaceString: paceString,
		Source:     domain.Strava,
		SourceId:   activity.Id,
		SportType:  sportType,
		TotalTime:  activity.ElapsedTime,
	}
}

func calculatePace(movingTime float64, distance float64) (pace int, paceString string) {
	distanceInKm := distance / kilometer
	secondsPerKm := movingTime / distanceInKm
	minutes := int(math.Floor(secondsPerKm / 60))
	seconds := int(secondsPerKm) % 60
	pace = int(secondsPerKm)
	paceString = fmt.Sprintf("%d:%02d", minutes, seconds)
	return
}

func FromStravaSportType(sportType SportType) domain.SportType {
	switch sportType {
	case SportTypeRun, SportTypeTrailRun:
		return domain.STRun
	case SportTypeRide, SportTypeGravelRide, SportTypeMountainBikeRide:
		return domain.STRide
	case SportTypeNordicSki:
		return domain.STXCSki
	case SportTypeRollerSki:
		return domain.STRollerSkis
	default:
		return domain.STOther
	}
}
