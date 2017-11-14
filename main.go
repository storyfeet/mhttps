package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"path"
)

var certlocs []string = []string{
	"/etc/letsencrypt/live/storyfeet.com/",
	"/etc/letsencrypt/live/www.enchantedconvoy.com/"}

func main() {
	tconf := &tls.Config{}

	for _, v := range certlocs {
		cert, err := tls.LoadX509KeyPair(path.Join(v, "fullchain.pem"), path.Join(v, "privkey.pem"))
		if err != nil {
			log.Fatal(err)
		}
		tconf.Certificates = append(tconf.Certificates, cert)
	}

	tconf.BuildNameToCertificate()

	go func() {
		log.Fatal(http.ListenAndServe(":8080", IHandler{}))
	}()

	sserv := http.Server{
		Addr:      ":8181",
		Handler:   SHandler{},
		TLSConfig: tconf,
	}

	log.Fatal(sserv.ListenAndServeTLS("", ""))

}
