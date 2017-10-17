package steamwebapi_test

import (
	"fmt"
	"testing"

	. "github.com/henkman/steamwebapi"
)

func TestGetServerInfo(t *testing.T) {
	si, err := GetServerInfo()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	fmt.Println(si)
}

func TestGetSupportedAPIList(t *testing.T) {
	key, err := getKey()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	ss, err := GetSupportedAPIList(key)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	for _, s := range ss {
		fmt.Printf("%+v\n", s)
	}
}
