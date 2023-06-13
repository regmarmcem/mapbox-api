package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Feature struct {
	ID         string   `json:"id"`
	Properties []string `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

type Geometry struct {
	Coordinates [][]int `json:"coordinates"`
	Type        string  `json:"type"`
}

type FeatureController struct {
	db *sql.DB
}

func NewFeatureController(db *sql.DB) *FeatureController {
	return &FeatureController{db: db}
}

func (c *FeatureController) PostFeature(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	json.NewEncoder(w).Encode(body)
}
