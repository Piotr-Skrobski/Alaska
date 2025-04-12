package models

type BoxOffice struct {
	Value    float64 `bson:"value" json:"value"`
	Currency string  `bson:"currency" json:"currency"`
}
