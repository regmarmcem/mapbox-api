package models

type Feature struct {
	ID         string   `json:"id"`
	Properties Property `json:"properties"`
	Geometry   Geometry `json:"geometry"`
}

type Geometry struct {
	Coordinates [][][]float64 `json:"coordinates"`
	Type        string  `json:"type"`
}

type Property struct {
	Title string `json:"title,omitempty"`
}