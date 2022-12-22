package entity

// Elasticsearch holds the configuration for the Elasticsearch client.
type Elasticsearch struct {
	Url       string    `yaml:"url" json:"url"`
	Username  string    `yaml:"username" json:"username"`
	Password  string    `yaml:"password" json:"password"`
	Highlight Highlight `yaml:"highlight" json:"highlight"`
}

type Highlight struct {
	PreTags  string `yaml:"preTags" json:"preTags"`
	PostTags string `yaml:"postTags" json:"postTags"`
}
