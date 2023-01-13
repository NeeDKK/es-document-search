package config

import (
	"encoding/json"
	"github.com/NeeDKK/esDocumentSearch/entity"
	"io/ioutil"
	"log"
)

var UnversityMap = make(map[string]string)

func ReadFileForMap() {

	data, err := ioutil.ReadFile("university.json")
	if nil != err {
		log.Fatalln("ReadFile ERROR:", err)
		return
	}
	var UnversityJson entity.UnversityJson
	err = json.Unmarshal(data, &UnversityJson)
	if nil != err {
		log.Fatalln("Unmarshal ERROR:", err)
		return
	}
	for _, v := range UnversityJson.RECORDS {
		UnversityMap[v.SchoolName] = "school"
	}
}
