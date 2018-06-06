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

	for {
		// Receive next message
		msg, err := receiver.Receive(ctx)
		if err != nil {
			log.Print("[!] reading message from AMQP:", err)
		}

		// Accept message
		msg.Accept()

		log.Printf("[*] message received on %q on %q: %s", os.Getenv("TYPE_OF_AMQP_USER"), os.Getenv("POD_NAME"), msg.GetData())
	}
}
