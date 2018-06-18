package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"qpid.apache.org/amqp"
	"qpid.apache.org/electron"
)

func main() {
	r := newReceiver()

	alternate := -1
	reject := os.Getenv("REJECT_ALTERNATE")
	for {
		if rm, err := r.Receive(); err == nil {
			if reject == "true" {
				alternate++
				if alternate%3 == 0 {
					rm.Accept()
				} else {
					rm.Release()
					log.Println("[!] msg was not accepted this time")
					continue
				}
			} else {
				rm.Accept()
			}

			log.Printf("[*] message received on %q on %q: %s", os.Getenv("TYPE_OF_AMQP_USER"), os.Getenv("POD_NAME"), rm.Message.Body())
		} else if fmt.Sprintf("%s", err) == "amqp:connection:forced: " {
			// this means we need to create receiver again
			r = newReceiver()
			continue
		} else if err == electron.Closed {
			break
		} else {
			log.Print("[!] reading message from AMQP:", err)
		}
	}
}

func newReceiver() electron.Receiver {
	sleepingTime := 1
	// persistent receiver creation
	for {
		log.Printf("sleeping while creating sender for %ds", sleepingTime)
		time.Sleep(time.Duration(sleepingTime) * time.Second)
		sleepingTime *= 2

		container := electron.NewContainer(os.Getenv("POD_NAME"))
		u, err := amqp.ParseURL(os.Getenv("AMQ_SERVER"))
		if err != nil {
			log.Printf("[!] parsing amqp url: %v", err)
			continue
		}
		addr := strings.TrimPrefix(u.Path, "/")

		c, err := container.Dial("tcp", u.Host)
		if err != nil {
			log.Printf("[!] dialing the amqp server: %v", err)
			continue
		}

		opts := []electron.LinkOption{
			electron.Source(addr),
			// electron.DurableSubscription(os.Getenv("TYPE_OF_AMQP_USER")),
		}

		r, err := c.Receiver(opts...)
		if err != nil {
			log.Printf("[!] creating receiver: %v", err)
			continue
		}
		return r
	}
}
