package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main(){
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Handling incoming orders.")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Handling users...")
	})

	port := 3000

	//load the TLS certificate and key
	cert := "cert.pem"
	key := "key.pem"

	//configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	//create a custom server
	server := &http.Server{
		Addr : fmt.Sprintf(":%d", port),
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	//enable http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Printf("Server is running on port: %d\n", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("error stating the server", err)
	}

	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalln("error stating the server")
	// }
	//Generating self signed certificates on our local computer using openssl:
	//openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
	//PEM - Privacy Enhanced Mail
	
	//key.pem = private key & cert.pem = public key
	//Generating these two keys separately.
	//First generate private key
	//openssl genpkey -algorithm RSA -out server.key -pkeyopt rsa_keygen_bits:2048
	//Now generate public key
	//openssl req -new -x509 -key server.key -out server.cert -days 365
}