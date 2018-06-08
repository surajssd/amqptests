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

func getInsecureSkipVerify() bool {
	insecureSkipVerify := os.Getenv("INSECURE_SKIP_VERIFY")
	if insecureSkipVerify == "true" {
		return true
	}
	return false
}

func isAnonAuth() bool {
	isAnon := os.Getenv("ANONYMOUS_AUTH")
	if isAnon == "true" {
		return true
	}
	return false
}

func GetAMQPSession() (*amqp.Session, func()) {

	username, password := getCreds()
	var options []amqp.ConnOption
	if isAnonAuth() {
		options = append(options, amqp.ConnSASLAnonymous())
	} else {
		options = append(options, amqp.ConnSASLPlain(username, password))
	}

	if getInsecureSkipVerify() {
		options = append(options, amqp.ConnTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}))
	}

	log.Println("[*] dialing AMQP server")
	client, err := amqp.Dial(
		os.Getenv("AMQ_SERVER"),
		options...,
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
