package main

import (
	"log"
	"os"
	"strings"

	"qpid.apache.org/amqp"
	"qpid.apache.org/electron"
)

func main() {

	container := electron.NewContainer(os.Getenv("POD_NAME"))

	url, err := amqp.ParseURL(os.Getenv("AMQ_SERVER"))
	if err != nil {
		log.Fatal("[!] parsing amqp url:", err)
	}

	c, err := container.Dial("tcp", url.Host)
	if err != nil {
		log.Fatal("[!] dialing the amqp server:", err)
	}
	addr := strings.TrimPrefix(url.Path, "/")
	opts := []electron.LinkOption{electron.Source(addr)}

	r, err := c.Receiver(opts...)
	if err != nil {
		log.Fatal("[!] creating receiver:", err)
	}

	for {
		// var msg []byte
		if rm, err := r.Receive(); err == nil {
			rm.Accept()
			// err := rm.Message.Decode(msg)
			// if err != nil {
			// 	log.Print("[!] decoding message from AMQP:", err)
			// 	continue
			// }
			log.Printf("[*] message received on %q on %q: %s", os.Getenv("TYPE_OF_AMQP_USER"), os.Getenv("POD_NAME"), rm.Message.Body())
		} else if err == electron.Closed {
			break
		} else {
			log.Print("[!] reading message from AMQP:", err)
		}
	}
}
