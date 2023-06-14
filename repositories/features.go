package repositories

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/regmarmcem/mapbox-api/models"
)

func SelectFeatureList(db *sql.DB) ([]models.Feature, error) {
	sqlStr := `
		select
			id, geometry
		from
			features
			;
	`
	rows, err := db.Query(sqlStr)
	if err != nil {
		log.Printf("SelectFeatureList failed: %s\n", err)
		return nil, err
	}
	defer rows.Close()

	featureArray := make([]models.Feature, 0)
	for rows.Next() {
		var feature models.Feature
		b := make([]byte, 0)
		rows.Scan(&feature.ID, &b)

		if err := json.Unmarshal(b, &feature.Geometry); err != nil {
			log.Panic(err)
		}

		feature.Type = "Feature"
		feature.Properties = make(map[string]string)
		featureArray = append(featureArray, feature)
	}
	return featureArray, nil
}

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
