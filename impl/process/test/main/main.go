package main

import (
	"fmt"
	"log"
)

func main() {
	r := GetWorldService()
	fmt.Print("ddd:", r)
	var m map[int]int
	m = make(map[int]int)
	for i := 0; i < 10; i++ {
		m[i] = i
	}
	for i := 0; i < 100; i++ {
		for k, _ := range m {
			log.Printf("k:%v", k)
			break
		}
	}
}

func GetWorldService() (r int) {
	defer func() {
		//world 可能在其他goroutine中已被移除.导致slice越界.
		if err := recover(); err != nil {
			log.Printf("GetWorldService panic:%v", err)
			r = 999
		}
	}()
	r = 5
	log.Panicf("dsfdsfds")
	return
}
