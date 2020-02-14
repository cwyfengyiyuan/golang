package model

import (
	"encoding/json"
	"log"
)

type Profile struct {
	Name string `json:"name"`
	Gender string `json:"gender"`
	User []string `json:"user"`
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	log.Printf("from: %s", s)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
//type Profile struct {
//	Name string
//	Gender string
//	Age int
//	Height int
//	Weight int
//	Income string
//	Marriage string
//	Education string
//	Occupation string
//	Hokou string
//	Xinzuo string
//	House string
//	Car string
//}