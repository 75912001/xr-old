package main

import "log"

func main() {
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
