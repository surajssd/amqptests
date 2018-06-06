package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/surajssd/amqptests/commons"
	"pack.ag/amqp"
)

func main() {
	session, closer := commons.GetAMQPSession()
	defer closer()

	ctx := context.Background()

	// This is a workaround to create multiple senders which is one sender
	// per address, which is provided comma separated, ideally the lib takes
	// care of it, but with current installation of ActiveMQ it does not work
	// so sender per address is the strategy being implemented
	addresses := strings.Split(os.Getenv("AMQP_ADDRESS"), ",")
	// create a senders
	var senders []*amqp.Sender

	for _, addr := range addresses {
		sender, err := session.NewSender(
			amqp.LinkTargetAddress(addr),
		)
		if err != nil {
			log.Fatal("[!] creating sender link:", err)
		}
		senders = append(senders, sender)
	}

	defer func() {
		for _, sender := range senders {
			sender.Close(ctx)
		}
	}()

	count := 0
	// endless for loop which keeps sending data
	for {
		msg := fmt.Sprintf("hello from %q on %q! %d", os.Getenv("TYPE_OF_AMQP_USER"), os.Getenv("POD_NAME"), count)

		for _, sender := range senders {
			err := sender.Send(ctx, amqp.NewMessage([]byte(msg)))
			if err != nil {
				log.Print("[!] error: sending message:", err)
			}
		}
		log.Println("[*] sent message:", msg)
		count++
		time.Sleep(2 * time.Second)
	}

}

/*
# when running locally without any openshift env

export AMQP_USERNAME=deployments && export AMQP_PASSWORD=test && export TYPE_OF_AMQP_USER=sender && export POD_NAME=host
export AMQP_USERNAME=deployments && export AMQP_PASSWORD=test && export TYPE_OF_AMQP_USER=receiver && export POD_NAME=host
*/
