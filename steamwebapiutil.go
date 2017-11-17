package steamwebapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type ServerInfo struct {
	Servertime       int    `json:"servertime"`
	Servertimestring string `json:"servertimestring"`
}

func GetServerInfo() (ServerInfo, error) {
	var si ServerInfo
	cli := http.Client{
		Timeout: time.Second * 10,
	}
	res, err := cli.Get(ENDPOINT + "/ISteamWebAPIUtil/GetServerInfo/v1/")
	if err != nil {
		return si, err
	}
	err = json.NewDecoder(res.Body).Decode(&si)
	res.Body.Close()
	if err != nil {
		return si, err
	}
	return si, nil
}

type Interface struct {
	Name    string `json:"name"`
	Methods []struct {
		Name       string `json:"name"`
		Version    int    `json:"version"`
		Httpmethod string `json:"httpmethod"`
		Parameters []struct {
			Name        string `json:"name"`
			Type        string `json:"type"`
			Optional    bool   `json:"optional"`
			Description string `json:"description"`
		} `json:"parameters"`
	} `json:"methods"`
}

func GetSupportedAPIList(key string) ([]Interface, error) {
	const URL = ENDPOINT + "/ISteamWebAPIUtil/GetSupportedAPIList/v1/"
	cli := http.Client{
		Timeout: time.Second * 10,
	}
	res, err := cli.Get(URL + "?key=" + key)
	if err != nil {
		return nil, err
	}
	var result struct {
		Apilist struct {
			Interfaces []Interface `json:"interfaces"`
		} `json:"apilist"`
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return result.Apilist.Interfaces, nil
}
