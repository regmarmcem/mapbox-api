package models

type Feature struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Properties map[string]string `json:"properties"`
	Geometry   Geometry          `json:"geometry"`
}

type Geometry struct {
	Coordinates [][][]float64 `json:"coordinates"`
	Type        string        `json:"type"`
}
