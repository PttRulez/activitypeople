package domain

import (
	"time"
)

type ActivityInfo struct {
	// The activity's distance, in meters
	Distance    float64 `json:"distance"`
	ElapsedTime int     `json:"elapsed_time"`
	Id          int     `json:"id"`

	// The activity's moving time, in seconds
	MovingTime int `json:"moving_time"`

	// The name of the activity
	Name        string    `json:"name"`
	PhotoCount  int       `json:"photo_count"`
	StartDate   time.Time `json:"start_date"`
	SportType   string    `json:"type"`
	WorkoutType int       `json:"workout_type"`
}

type SportType string

const (

	// SportTypeAlpineSki captures enum value "AlpineSki"
	SportTypeAlpineSki SportType = "AlpineSki"

	// SportTypeBackcountrySki captures enum value "BackcountrySki"
	SportTypeBackcountrySki SportType = "BackcountrySki"

	// SportTypeBadminton captures enum value "Badminton"
	SportTypeBadminton SportType = "Badminton"

	// SportTypeCanoeing captures enum value "Canoeing"
	SportTypeCanoeing SportType = "Canoeing"

	// SportTypeCrossfit captures enum value "Crossfit"
	SportTypeCrossfit SportType = "Crossfit"

	// SportTypeEBikeRide captures enum value "EBikeRide"
	SportTypeEBikeRide SportType = "EBikeRide"

	// SportTypeElliptical captures enum value "Elliptical"
	SportTypeElliptical SportType = "Elliptical"

	// SportTypeEMountainBikeRide captures enum value "EMountainBikeRide"
	SportTypeEMountainBikeRide SportType = "EMountainBikeRide"

	// SportTypeGolf captures enum value "Golf"
	SportTypeGolf SportType = "Golf"

	// SportTypeGravelRide captures enum value "GravelRide"
	SportTypeGravelRide SportType = "GravelRide"

	// SportTypeHandcycle captures enum value "Handcycle"
	SportTypeHandcycle SportType = "Handcycle"

	// SportTypeHighIntensityIntervalTraining captures enum value "HighIntensityIntervalTraining"
	SportTypeHighIntensityIntervalTraining SportType = "HighIntensityIntervalTraining"

	// SportTypeHike captures enum value "Hike"
	SportTypeHike SportType = "Hike"

	// SportTypeIceSkate captures enum value "IceSkate"
	SportTypeIceSkate SportType = "IceSkate"

	// SportTypeInlineSkate captures enum value "InlineSkate"
	SportTypeInlineSkate SportType = "InlineSkate"

	// SportTypeKayaking captures enum value "Kayaking"
	SportTypeKayaking SportType = "Kayaking"

	// SportTypeKitesurf captures enum value "Kitesurf"
	SportTypeKitesurf SportType = "Kitesurf"

	// SportTypeMountainBikeRide captures enum value "MountainBikeRide"
	SportTypeMountainBikeRide SportType = "MountainBikeRide"

	// SportTypeNordicSki captures enum value "NordicSki"
	SportTypeNordicSki SportType = "NordicSki"

	// SportTypePickleball captures enum value "Pickleball"
	SportTypePickleball SportType = "Pickleball"

	// SportTypePilates captures enum value "Pilates"
	SportTypePilates SportType = "Pilates"

	// SportTypeRacquetball captures enum value "Racquetball"
	SportTypeRacquetball SportType = "Racquetball"

	// SportTypeRide captures enum value "Ride"
	SportTypeRide SportType = "Ride"

	// SportTypeRockClimbing captures enum value "RockClimbing"
	SportTypeRockClimbing SportType = "RockClimbing"

	// SportTypeRollerSki captures enum value "RollerSki"
	SportTypeRollerSki SportType = "RollerSki"

	// SportTypeRowing captures enum value "Rowing"
	SportTypeRowing SportType = "Rowing"

	// SportTypeRun captures enum value "Run"
	SportTypeRun SportType = "Run"

	// SportTypeSail captures enum value "Sail"
	SportTypeSail SportType = "Sail"

	// SportTypeSkateboard captures enum value "Skateboard"
	SportTypeSkateboard SportType = "Skateboard"

	// SportTypeSnowboard captures enum value "Snowboard"
	SportTypeSnowboard SportType = "Snowboard"

	// SportTypeSnowshoe captures enum value "Snowshoe"
	SportTypeSnowshoe SportType = "Snowshoe"

	// SportTypeSoccer captures enum value "Soccer"
	SportTypeSoccer SportType = "Soccer"

	// SportTypeSquash captures enum value "Squash"
	SportTypeSquash SportType = "Squash"

	// SportTypeStairStepper captures enum value "StairStepper"
	SportTypeStairStepper SportType = "StairStepper"

	// SportTypeStandUpPaddling captures enum value "StandUpPaddling"
	SportTypeStandUpPaddling SportType = "StandUpPaddling"

	// SportTypeSurfing captures enum value "Surfing"
	SportTypeSurfing SportType = "Surfing"

	// SportTypeSwim captures enum value "Swim"
	SportTypeSwim SportType = "Swim"

	// SportTypeTableTennis captures enum value "TableTennis"
	SportTypeTableTennis SportType = "TableTennis"

	// SportTypeTennis captures enum value "Tennis"
	SportTypeTennis SportType = "Tennis"

	// SportTypeTrailRun captures enum value "TrailRun"
	SportTypeTrailRun SportType = "TrailRun"

	// SportTypeVelomobile captures enum value "Velomobile"
	SportTypeVelomobile SportType = "Velomobile"

	// SportTypeVirtualRide captures enum value "VirtualRide"
	SportTypeVirtualRide SportType = "VirtualRide"

	// SportTypeVirtualRow captures enum value "VirtualRow"
	SportTypeVirtualRow SportType = "VirtualRow"

	// SportTypeVirtualRun captures enum value "VirtualRun"
	SportTypeVirtualRun SportType = "VirtualRun"

	// SportTypeWalk captures enum value "Walk"
	SportTypeWalk SportType = "Walk"

	// SportTypeWeightTraining captures enum value "WeightTraining"
	SportTypeWeightTraining SportType = "WeightTraining"

	// SportTypeWheelchair captures enum value "Wheelchair"
	SportTypeWheelchair SportType = "Wheelchair"

	// SportTypeWindsurf captures enum value "Windsurf"
	SportTypeWindsurf SportType = "Windsurf"

	// SportTypeWorkout captures enum value "Workout"
	SportTypeWorkout SportType = "Workout"

	// SportTypeYoga captures enum value "Yoga"
	SportTypeYoga SportType = "Yoga"
)
