package model

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var sm sync.Map
	sm.Store("aa", "bb")
	sm.Store("bb", []string{"11"})
	sm.Store("bb", []string{"22"})
	sm.Store("cc", []string{"caca"})
	cc := "lalalal"
	sm.Store("cc", func(cc string) interface{} {
		data, _ := sm.Load("cc")
		data = append(data.([]string), cc)
		//data.([]string) = append(data.([]string), cc)
		return data
	}(cc))
	t.Log(sm.Load("bb"))
	sm.Range(func(key, value interface{}) bool {
		t.Log("key:", key, "value:", value)
		return true
	})
}
