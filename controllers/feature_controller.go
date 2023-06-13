package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/regmarmcem/mapbox-api/models"
)


type FeatureController struct {
	db *sql.DB
}

func NewFeatureController(db *sql.DB) *FeatureController {
	return &FeatureController{db: db}
}

func (c *FeatureController) PostFeature(w http.ResponseWriter, r *http.Request) {
	var reqFeature models.Feature

	if err := json.NewDecoder(r.Body).Decode(&reqFeature); err != nil {
		http.Error(w, "Cannot decode request body", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reqFeature)
}
