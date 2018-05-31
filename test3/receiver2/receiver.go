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

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// num := r.Intn(99999)
	// numStr := fmt.Sprintf("%d", num)

	// create a receiver
	receiver, err := session.NewReceiver(
		amqp.LinkSourceAddress("osio.space"),
		// amqp.LinkProperty("key"+numStr, "value"+numStr),
		// amqp.LinkSessionFilter("foobar"),
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
