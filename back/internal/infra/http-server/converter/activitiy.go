package converter

import (
	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/http-server/contracts"
)

func FromActivityToActivityResponse(a domain.Activity,
) contracts.ActivityResponse {
	return contracts.ActivityResponse{
		Calories:    a.Calories,
		Description: a.Description,
		Distance:    a.Distance,
		Date:        a.Date,
		Elevate:     a.Elevate,
		Heartrate:   a.Heartrate,
		Id:          a.Id,
		Name:        a.Name,
		Pace:        a.Pace,
		PaceString:  a.PaceString,
		Source:      a.Source,
		SourceId:    a.SourceId,
		SportType:   a.SportType,
		TotalTime:   a.TotalTime,
	}
}
