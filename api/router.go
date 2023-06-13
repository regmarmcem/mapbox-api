package api

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/regmarmcem/mapbox-api/controllers"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()
	c := controllers.NewFeatureController(db)
	r.MethodFunc(http.MethodPost, "/feature", c.PostFeature)
	return r
}
