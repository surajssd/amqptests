package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/surajssd/amqptests/commons"
	"pack.ag/amqp"
)

func main() {
	session, closer := commons.GetAMQPSession()
	defer closer()

	ctx := context.Background()
	// create a sender
	sender, err := session.NewSender(
		amqp.LinkTargetAddress(os.Getenv("AMQP_ADDRESS")),
	)
	if err != nil {
		log.Fatal("[!] creating sender link:", err)
	}
	defer sender.Close(ctx)

	count := 0
	// endless for loop which keeps sending data
	for {
		msg := fmt.Sprintf("hello from %q on %q! %d", os.Getenv("TYPE_OF_AMQP_USER"), os.Getenv("POD_NAME"), count)
		err = sender.Send(ctx, amqp.NewMessage([]byte(msg)))
		if err != nil {
			log.Print("[!] error: sending message:", err)
		}
		log.Println("[*] sent message:", msg)
		count++
		time.Sleep(2 * time.Second)
	}

}
