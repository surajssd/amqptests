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

	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// num := r.Intn(99999)
	// numStr := fmt.Sprintf("%d", num)
	// numStr := "2"

	username, password := getCreds()
	client, err := amqp.Dial(
		"amqps://messaging-maas-aslakzredhatzcom.6a63.fuse-ignite.openshiftapps.com:443",
		amqp.ConnSASLPlain(username, password),
		// amqp.ConnProperty("key"+numStr, "value"+numStr),
		amqp.ConnTLSConfig(&tls.Config{
			InsecureSkipVerify: true,
		}),
	)
	if err != nil {
		log.Fatal("[!] dialing AMQP server:", err)
	}

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("[!] creating AMQP session:", err)
	}

	return session, func() {
		client.Close()
	}
}
