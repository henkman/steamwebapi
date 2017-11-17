package steamwebapi_test

import (
	"fmt"
	"testing"

	. "github.com/henkman/steamwebapi"
)

func TestGetFriendList(t *testing.T) {
	key, err := getKey()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	steamid, err := getSteamId()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	ss, err := GetFriendList(key, steamid)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	for _, s := range ss {
		fmt.Printf("%+v\n", s)
	}
}

func TestGetPlayerSummaries(t *testing.T) {
	key, err := getKey()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	steamid, err := getSteamId()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	ss, err := GetPlayerSummaries(key, []string{steamid})
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	for _, s := range ss {
		fmt.Printf("%+v\n", s)
	}
}
