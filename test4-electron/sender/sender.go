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
	s, err := c.Sender(electron.Target(addr))
	if err != nil {
		log.Fatal("[!] creating sender:", err)
	}

	count := 0
	// endless for loop which keeps sending data
	for {
		m := amqp.NewMessage()
		msg := fmt.Sprintf("hello from %q on %q! %d", os.Getenv("TYPE_OF_AMQP_USER"), os.Getenv("POD_NAME"), count)
		m.Marshal(msg)
		m.SetFirstAcquirer(true)
		outcome := s.SendSync(m)
		if outcome.Error != nil {
			log.Print("[!] sending message:", outcome.Value, ", error:", outcome.Error)
			// continue to retry sending the message
			continue
		} else if outcome.Status != electron.Accepted {
			log.Print("[!] sending message:", outcome.Value, ", unexpected status:", outcome.Status)
			// continue to retry sending the message
			continue
		}

		log.Println("[*] sent message:", msg)
		count++
		time.Sleep(2 * time.Second)
	}

}
