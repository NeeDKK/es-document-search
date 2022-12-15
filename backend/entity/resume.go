package entity

type Resume struct {
	BASE_MODEL
	Name    string `json:"name"`
	School  string `json:"school"`
	Content string `json:"content"`
}
