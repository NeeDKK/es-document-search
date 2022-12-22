package entity

type Config struct {
	Elasticsearch Elasticsearch `yaml:"elasticsearch" json:"elasticsearch"`
	Server        Server        `yaml:"server" json:"server"`
}
