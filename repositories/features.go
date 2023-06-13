package repositories

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/regmarmcem/mapbox-api/models"
)

func InsertFeature(db *sql.DB, feature models.Feature) (models.Feature, error) {
	const sqlStr = `
		insert into features (id, geometry) values ($1, $2);
	`
	
	var newFeature models.Feature
	newFeature.ID, newFeature.Geometry = feature.ID, feature.Geometry

	geometry_json, _ := json.Marshal(feature.Geometry)
	
	_, err := db.Exec(sqlStr, newFeature.ID, geometry_json)
	if err != nil {
		log.Println("InsertFeature failed")
		return models.Feature{}, err
	}
	
	return newFeature, nil
}