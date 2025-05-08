package main

import (
	"ToDo_List/router"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("App started")
	setLogFile()
	r := router.Router()
	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create the TLS Config with the CA pool and enable Client certificate validation
	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	// Create a Server instance to listen on port 8443 with the TLS config
	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
		Handler:   r,
	}
	log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}

func setLogFile() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panicf("Failed to open file")
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Log file created!")
}
