package crawler

import (
	"encoding/json"
	"log"
)

type Data struct {

	Title string
	URL string
	Phrase []string
}


func (d *Data) ToJson() string {

	b, err := json.Marshal(d)

	if err != nil {
		log.Printf("Unable to convert %v to json", d)
		return ""
	}
	return string(b)
}



func ParseJsonArray(s string, datas *[]Data) error {
	err := json.Unmarshal([]byte(s), datas)
	if err != nil {
		return err
	}
	return nil
}


func ToJsonArray(data []Data) string {

	b, err := json.Marshal(data)

	if err != nil {
		log.Printf("Unable to convert %v to json", data)
		return ""
	}

	return string(b)
}