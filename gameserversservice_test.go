package steamwebapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetServerList(t *testing.T) {
	var key string
	{
		kf, err := ioutil.ReadFile("testkey.json")
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		if err := json.Unmarshal(kf, &key); err != nil {
			t.Error(err)
			t.Fail()
			return
		}
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
