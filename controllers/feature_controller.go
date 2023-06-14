package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/regmarmcem/mapbox-api/models"
	"github.com/regmarmcem/mapbox-api/repositories"
)

type FeatureController struct {
	db *sql.DB
}

func NewFeatureController(db *sql.DB) *FeatureController {
	return &FeatureController{db: db}
}

func (c *FeatureController) ListFeatures(w http.ResponseWriter, r *http.Request) {
	features, err := repositories.SelectFeatureList(c.db)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(features)
}

func (c *FeatureController) PostFeature(w http.ResponseWriter, r *http.Request) {
	var reqFeature models.Feature

	if err := json.NewDecoder(r.Body).Decode(&reqFeature); err != nil {
		http.Error(w, "Cannot decode request body", http.StatusInternalServerError)
		return
	}
	feature, err := repositories.InsertFeature(c.db, reqFeature)
	if err != nil {
		http.Error(w, "Cannot insert specified feature", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(feature)
}
