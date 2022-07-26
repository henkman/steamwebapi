package steamwebapi

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"
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
	res, err := httpGetWithTimeout(URL+"?"+ps.Encode(), time.Second*10)
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

type PlayerSummary struct {
	Steamid                  string `json:"steamid"`
	Communityvisibilitystate int    `json:"communityvisibilitystate"`
	Profilestate             int    `json:"profilestate"`
	Personaname              string `json:"personaname"`
	Commentpermission        int    `json:"commentpermission,omitempty"`
	Profileurl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	Avatarmedium             string `json:"avatarmedium"`
	Avatarfull               string `json:"avatarfull"`
	Avatarhash               string `json:"avatarhash"`
	Lastlogoff               int    `json:"lastlogoff"`
	Personastate             int    `json:"personastate"`
	Realname                 string `json:"realname,omitempty"`
	Primaryclanid            string `json:"primaryclanid,omitempty"`
	Timecreated              int    `json:"timecreated,omitempty"`
	Personastateflags        int    `json:"personastateflags,omitempty"`
	Gameserverip             string `json:"gameserverip,omitempty"`
	Gameserversteamid        string `json:"gameserversteamid,omitempty"`
	Gameextrainfo            string `json:"gameextrainfo,omitempty"`
	Gameid                   string `json:"gameid,omitempty"`
	Lobbysteamid             string `json:"lobbysteamid,omitempty"`
	Loccountrycode           string `json:"loccountrycode,omitempty"`
	Locstatecode             string `json:"locstatecode,omitempty"`
	Loccityid                int    `json:"loccityid,omitempty"`
}

func GetPlayerSummaries(key string, steamids []string) ([]PlayerSummary, error) {
	const URL = ENDPOINT + "/ISteamUser/GetPlayerSummaries/v0002/"
	ps := url.Values{
		"key":      []string{key},
		"steamids": []string{strings.Join(steamids, ",")},
	}
	res, err := httpGetWithTimeout(URL+"?"+ps.Encode(), time.Second*10)
	if err != nil {
		return nil, err
	}
	var result struct {
		Response struct {
			Players []PlayerSummary `json:"players"`
		} `json:"response"`
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return result.Response.Players, nil
}
