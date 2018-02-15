package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"

	"github.com/drhayt/jwtclient"
)

func main() {

	var (
		jwtcert = flag.String("jwtcert", "jwt.crt", "The cert file to use")
		url     = flag.String("url", "https://authentication.sgtec.io", "The url to use to authenticate.")
	)

	flag.Parse()

	// get the validation cert.

	pemFile, err := os.Open(*jwtcert)
	if err != nil {
		panic(err)
	}

	pemBytes := bytes.NewBuffer(nil)

	_, err = io.Copy(pemBytes, pemFile)
	if err != nil {
		panic(err)
	}

	config := jwtclient.Config{
		AuthKey:    "admin",
		AuthSecret: "admin",
		URL:        *url,
		Insecure:   true,
	}

	client, err := jwtclient.New(&config)
	if err != nil {
		log.Fatalf("Unable to create client: %s", err)
	}

	// for i := 0; i < 10000; i++ {
	for i := 0; i < 16; i++ {
		foo, err := client.RetrieveToken()
		if err != nil {
			log.Printf("Unable to retrieve token: %s", err)
		}
		log.Printf("Token is %s", foo)
		// time.Sleep(1 * time.Second)
		// fmt.Printf("Token is %s\nSleeping 10 seconds\n\n", foo)
		// time.Sleep(10 * time.Second)
	}

}
