package main

import (
  "os"
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go/http3"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
  w, _ := os.Create("./keylog.log")

	r := http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS13,
			MaxVersion: tls.VersionTLS13,
      KeyLogWriter: w,
		},
	}
  req, _ := http.NewRequest("GET", "https://localhost:18443", nil)

	resp, err := r.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))

}
