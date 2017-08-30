package steamwebapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Server struct {
	Addr       string `json:"addr"`
	Gameport   int    `json:"gameport"`
	Steamid    string `json:"steamid"`
	Name       string `json:"name"`
	Appid      int    `json:"appid"`
	Gamedir    string `json:"gamedir"`
	Version    string `json:"version"`
	Product    string `json:"product"`
	Region     int    `json:"region"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"max_players"`
	Bots       int    `json:"bots"`
	Map        string `json:"map"`
	Secure     bool   `json:"secure"`
	Dedicated  bool   `json:"dedicated"`
	Os         string `json:"os"`
	Gametype   string `json:"gametype"`
}

// https://developer.valvesoftware.com/wiki/Master_Server_Query_Protocol#Filter
func GetServerList(key string, limit uint, filter string) ([]Server, error) {
	const URL = ENDPOINT + "/IGameServersService/GetServerList/v1/"
	ps := url.Values{
		"filter": []string{filter},
		"limit":  []string{fmt.Sprint(limit)},
		"key":    []string{key},
	}
	res, err := http.Get(URL + "?" + ps.Encode())
	if err != nil {
		return nil, err
	}
	var result struct {
		Response struct {
			Servers []Server `json:"servers"`
		} `json:"response"`
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return result.Response.Servers, nil
}
