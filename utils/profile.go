package utils

import (
	"encoding/json"
	"io/ioutil"
)

type InstagramProfile struct {
	Username        string `json:"username"`
	URL             string `json:"url"`
	LastCheckedTime int64  `json:"lastCheckedTime"`
}

func ReadProfilesFromJSON(filename string) ([]InstagramProfile, error) {
	profilesBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var profiles []InstagramProfile
	err = json.Unmarshal(profilesBytes, &profiles)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}
