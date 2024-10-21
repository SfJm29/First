package main

import (
	"fmt"
	"time"
)

func gh(nd chan int, d chan bool) {
	for {
		//time.Sleep(time.Second * 2)
		select {
		case <-nd:
			fmt.Println("получил данные")
		case <-d:
			fmt.Println("end")
			return
		}
	}
}

func main() {
	//fmt.Println("asdf")
	dn := make(chan bool)
	ndd := make(chan int)
	go gh(ndd, dn)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 5)
		//ndd <- 3
		select {
		case _, ok := <-dn:
			if ok {
				fmt.Println("Канал закрыли")
				close(dn)
			}
		default:
			fmt.Println("Канал закрыт")
		}

	}
}
