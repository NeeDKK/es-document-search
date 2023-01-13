package entity

type UnversityJson struct {
	RECORDS []Unversity `json:"RECORDS"`
}

type Unversity struct {
	Id                 string      `json:"id"`
	SchoolName         string      `json:"school_name"`
	SchoolCode         string      `json:"school_code"`
	CompetentDarptment string      `json:"competent_darptment"`
	Location           string      `json:"location"`
	Level              string      `json:"level"`
	Type               interface{} `json:"type"`
}
