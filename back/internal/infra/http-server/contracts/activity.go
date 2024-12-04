package contracts

import (
	"time"

	"github.com/pttrulez/activitypeople/internal/domain"
)

type ActivityDayResponse struct {
	Date       string             `json:"date"`
	Activities []ActivityResponse `json:"activities"`
}

type ActivityResponse struct {
	Calories    int    `json:"calories"`
	Description string `json:"description"`

	// расстояние в километрах
	Distance int       `json:"distance"`
	Date     time.Time `json:"date"`

	// общий подъём в метрах
	Elevate int `json:"elevate"`

	// средний пульс
	Heartrate int    `json:"heartrate"`
	Id        int    `json:"id"`
	Name      string `json:"name"`

	// темп в секундах
	Pace       int                   `json:"pace"`
	PaceString string                `json:"paceString"`
	Source     domain.ActivitySource `json:"source"`
	SourceId   int64                 `json:"sourceId"`

	// айдишник активити у поставщика (гармин, страва и т.д.)
	SportType domain.SportType `json:"sportType"`

	// время в секундах
	TotalTime int `json:"totalTime"`
}
