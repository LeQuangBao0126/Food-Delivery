package main

import (
	"fmt"
	"time"
)

func main(){

	queue:= make(chan int , 5000)

	for i := 1 ; i <= 5000;i++{
		queue<- i
	}


	go func() {
			for {
				time.Sleep(time.Second)
				fmt.Println("worker 1", <-queue)
			}
	}()


	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("worker 2", <-queue)
		}
	}()




	time.Sleep(time.Second * 10)
	fmt.Println("End")
}


