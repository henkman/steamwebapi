package steamwebapi_test

import (
	"fmt"
	"testing"

	. "github.com/henkman/steamwebapi"
)

func TestGetOwnedGames(t *testing.T) {
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
	games, err := GetOwnedGames(key, steamid, []string{"418460"}, true)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	for _, game := range games {
		fmt.Printf("%+v\n", game)
	}
}
