package main

import (
	v2 "go-file-splitter/v2"
	"time"
)

func main() {
	//v1.Do()
	v2.Do()

	for true {
		time.Sleep(1 * time.Second)
	}
}