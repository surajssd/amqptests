package commons

import (
	"crypto/tls"
	"log"
	"os"

	"pack.ag/amqp"
)

func getCreds() (string, string) {
	username := os.Getenv("AMQP_USERNAME")
	password := os.Getenv("AMQP_PASSWORD")

	return username, password
}

func GetAMQPSession() (*amqp.Session, func()) {

	username, password := getCreds()
	client, err := amqp.Dial(
		os.Getenv("AMQ_SERVER"),
		amqp.ConnSASLPlain(username, password),
		amqp.ConnTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}),
	)
	if err != nil {
		log.Fatal("[!] dialing AMQP server:", err)
	}
	log.Println("[*] Authenticated with the bus.")

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("[!] creating AMQP session:", err)
	}
	log.Println("[*] Session created.")

	return session, func() {
		client.Close()
	}
}
