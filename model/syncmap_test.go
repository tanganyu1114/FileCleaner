package model

import (
	"fmt"
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

func TestNUmberAdd(t *testing.T) {
	defer func() {
		fmt.Println("defer")
	}()
	if true {
		fmt.Println("true")
		return
	}
	fmt.Println("aaaa")

}

func TestChtwiece(t *testing.T) {
	var testCH = make(chan int, 1)
	var sigCH = make(chan bool)
	var dataCH = make(chan bool)
	go func() {
		for {
			select {
			case <-sigCH:
				testCH <- 1
				testCH <- 2
			case data := <-dataCH:
				fmt.Println("data", data)
			}
		}
	}()
	sigCH <- true
	firstfmt(testCH)
	secondfmt(testCH)
	dataCH <- false
}
func firstfmt(testCH chan int) {
	aa := <-testCH
	fmt.Println("111", aa)
}
func secondfmt(testCH chan int) {
	aa := <-testCH
	fmt.Println("222", aa)
}

func TestMapString(t *testing.T) {
	var test = make(map[int][]int)
	test[5] = append(test[5], 1)
	t.Log(test)
	test[1] = append(test[1], 1111)
	test[1] = append(test[1], 1111)
	var test2 = make(map[int]int)
	test2[11] = 22
	test2[11] = 33
	test2[11] = 44
	fmt.Println(test2)
	test2[11] = 22
	fmt.Println(test2)
}
