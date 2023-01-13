package entity

type Source struct {
	Attachment Resume `json:"attachment"`
	School     string `json:"school"`
	Name       string `json:"name"`
	ID         int    `json:"ID"`
}
