package steamwebapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type Game struct {
	Appid                    int    `json:"appid"`
	Name                     string `json:"name"`
	Playtime2Weeks           int    `json:"playtime_2weeks"`
	PlaytimeForever          int    `json:"playtime_forever"`
	ImgIconURL               string `json:"img_icon_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
	PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
	PlaytimeMacForever       int    `json:"playtime_mac_forever"`
	PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
}

func GetOwnedGames(key, steamid string, appids_filter []string, include_appinfo bool) ([]Game, error) {
	const URL = ENDPOINT + "/IPlayerService/GetOwnedGames/v1/"
	ps := url.Values{
		"steamid": []string{steamid},
		"key":     []string{key},
	}
	if appids_filter != nil && len(appids_filter) > 0 {
		for i, af := range appids_filter {
			ps[fmt.Sprintf("appids_filter[%d]", i)] = []string{af}
		}
	}
	if include_appinfo {
		ps["include_appinfo"] = []string{"true"}
	}
	res, err := httpGetWithTimeout(URL+"?"+ps.Encode(), time.Second*10)
	if err != nil {
		return nil, err
	}
	var result struct {
		Response struct {
			GameCount int    `json:"game_count"`
			Games     []Game `json:"games"`
		} `json:"response"`
	}
	err = json.NewDecoder(res.Body).Decode(&result)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return result.Response.Games, nil
}
