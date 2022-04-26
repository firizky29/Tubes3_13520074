package transaction

type PredictionRequest struct {
	UserName          string `json:"UserName"`
	DNA               string `json:"DNA"`
	DiseasePrediction string `json:"DiseasePrediction"`
}
