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
	s := newSender()

	count := 0
	// endless for loop which keeps sending data
	for {
		m := amqp.NewMessage()
		msg := fmt.Sprintf("hello from %q on %q! %d", os.Getenv("TYPE_OF_AMQP_USER"), os.Getenv("POD_NAME"), count)
		m.Marshal(msg)
		outcome := s.SendSync(m)
		if fmt.Sprintf("%s", outcome.Error) == "amqp:connection:forced: " {
			// this means we need to create sender again
			s = newSender()
			continue
		} else if outcome.Error != nil {
			log.Print("[!] sending message:", outcome.Value, ", error:", outcome.Error)
			// continue to retry sending the message
			time.Sleep(time.Second)
			continue
		} else if outcome.Status != electron.Accepted {
			log.Print("[!] sending message:", outcome.Value, ", unexpected status:", outcome.Status)
			// continue to retry sending the message
			time.Sleep(time.Second)
			continue
		}

		log.Println("[*] sent message:", msg)
		count++
		time.Sleep(2 * time.Second)
	}

}

func newSender() electron.Sender {

	sleepingTime := 1
	// persistent sender creation
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
		s, err := c.Sender(electron.Target(addr))
		if err != nil {
			log.Printf("[!] creating sender: %v", err)
			continue
		}
		return s
	}
}
