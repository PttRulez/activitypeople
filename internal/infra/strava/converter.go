package strava

import (
	"github.com/pttrulez/activitypeople/internal/domain"
)

func FromStravaToActivity(activity ActivityResponse) domain.Activity {
	return domain.Activity{
		Distance:  activity.Distance,
		TotalTime: activity.ElapsedTime,
		Name:      activity.Name,
		SportType: FromStravaSportType(activity.SportType),
	}
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
