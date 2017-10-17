package steamwebapi_test

import (
	"encoding/json"
	"io/ioutil"
)

func getSteamId() (string, error) {
	var id string
	kf, err := ioutil.ReadFile("teststeamid.json")
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(kf, &id); err != nil {
		return "", err
	}
	return id, nil
}

func getKey() (string, error) {
	var key string
	kf, err := ioutil.ReadFile("testkey.json")
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(kf, &key); err != nil {
		return "", err
	}
	return key, nil
}
