package main

import (
	"context"
	"log"
	"os"

	"github.com/surajssd/amqptests/commons"
	"pack.ag/amqp"
)

func main() {
	session, closer := commons.GetAMQPSession()
	defer closer()

	ctx := context.Background()

	// create a receiver
	receiver, err := session.NewReceiver(
		amqp.LinkSourceAddress(os.Getenv("AMQP_ADDRESS")),
	)
	if err != nil {
		log.Fatal("[!] creating receiver link:", err)
	}
	defer receiver.Close(ctx)

	var alternate bool
	reject := os.Getenv("REJECT_ALTERNATE")
	for {
		// Receive next message
		msg, err := receiver.Receive(ctx)
		if err != nil {
			log.Print("[!] reading message from AMQP:", err)
		}

		// if the env var is set to reject the message then reject it
		if reject == "true" {
			// we will reject the alternate messages
			if alternate {
				// this API is used to reject the message
				msg.Release()
			} else {
				msg.Accept()
			}
			alternate = !alternate
		} else {
			msg.Accept()
		}

		// this will be printed no matter whether the message is rejected or not
		// and since we are using the alternate reject message strategy so the messages
		// will be seen twice
		log.Printf("[*] message received on %q on %q: %s", os.Getenv("TYPE_OF_AMQP_USER"), os.Getenv("POD_NAME"), msg.GetData())
	}
}
