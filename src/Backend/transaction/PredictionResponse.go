package transaction

type PredictionResponse struct {
	PredictionDate  string `json:"PredictionDate"`
	UserName        string `json:"UserName"`
	DiseaseName     string `json:"DiseaseName"`
	SimilarityLevel string `json:"Similarity"`
	Status          string `json:"Status"`
}
