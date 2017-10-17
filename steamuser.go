package steamwebapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Friend struct {
	Steamid      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

func GetFriendList(key, steamid string) ([]Friend, error) {
	const URL = ENDPOINT + "/ISteamUser/GetFriendList/v1/"
	ps := url.Values{
		"key":     []string{key},
		"steamid": []string{steamid},
	}
	res, err := http.Get(URL + "?" + ps.Encode())
	if err != nil {
		return nil, err
	}
	var result struct {
		Friendslist struct {
			Friends []Friend `json:"friends"`
		} `json:"friendslist"`
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return result.Friendslist.Friends, nil
}
