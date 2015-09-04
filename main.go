package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"willnorris.com/go/imageproxy"
)

func main() {

	addr := os.Getenv("ADDRESS")
	if addr == "" {
		log.Fatal("No address provided for the imageproxy")
	}

	p := imageproxy.NewProxy(nil, nil)

	if os.Getenv("WHITELIST") != "" {
		p.Whitelist = strings.Split(os.Getenv("WHITELIST"), ",")
	}

	if os.Getenv("BASEURL") != "" {
		var err error
		p.DefaultBaseURL, err = url.Parse(os.Getenv("BASEURL"))
		if err != nil {
			log.Fatalf("error parsing baseURL: %v", err)
		}
	}

	server := &http.Server{
		Addr:    addr,
		Handler: p,
	}

	fmt.Printf("imageproxy listening on " + addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
