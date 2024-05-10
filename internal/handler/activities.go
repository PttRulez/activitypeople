package handler

import (
	"antiscoof/internal/view/activities"
	"net/http"
)

func GetActivitiesPage(w http.ResponseWriter, r *http.Request) error {
	return activities.Index().Render(r.Context(), w)
}
