package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printData(data interface{}) {
	for i := 1; i <= 4; i++ {
		fmt.Println(data, i)
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	}
}

func main() {
	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	go printData(data1)
	go printData(data2)

	time.Sleep(1 * time.Second)
}
