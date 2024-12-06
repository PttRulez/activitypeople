package converter

import (
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
)

func FromUserSettingsRequestToUserSettings(req contracts.UserSettingsRequest) domain.UserSettings {
	return domain.UserSettings{
		BMR:                 req.BMR,
		CaloriesPer100Steps: req.CaloriesPer100Steps,
	}
}
