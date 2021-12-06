package main

import (
	"ProjectDelivery2/pubsub"
	"ProjectDelivery2/pubsub/pblocal"
	"context"
	"log"
	"time"
)

func main(){
	localPB := pblocal.NewPubSub()

	var topic pubsub.Topic = "Order Created"

	sub1 ,_ := localPB.Subscribe(context.Background(),topic)
	sub2 ,_ := localPB.Subscribe(context.Background(),topic)

	localPB.Publish(context.Background(),topic,pubsub.NewMessage("1"))

	go func(){
		for{
			log.Println("Con 1",(<-sub1).Data())
			time.Sleep(time.Second)
		}
	}()
	go func(){
		for{
			log.Println("Con 2",(<-sub2).Data())
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 4)
}
