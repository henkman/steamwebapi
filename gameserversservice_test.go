package steamwebapi_test

import (
	"fmt"
	"testing"

	. "github.com/henkman/steamwebapi"
)

func TestGetServerList(t *testing.T) {
	key, err := getKey()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	ss, err := GetServerList(key, 10, `\gamedir\RS2\name_match\*40-1*`)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	for _, s := range ss {
		fmt.Printf("%+v\n", s)
	}
}
