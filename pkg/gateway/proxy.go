package gateway

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type reverseProxy struct {
	httputil.ReverseProxy
	destination Destination
}

func NewReverseProxy(destination Destination) *reverseProxy {
	p := &reverseProxy{
		destination: destination,
	}
	p.Director = func(req *http.Request) {
		req.URL.Scheme = "http"
		host := fmt.Sprintf("%s:%d", p.destination.Host, p.destination.Port)
		req.URL.Host = host
		req.Host = host

		b, err := httputil.DumpRequest(req, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
		fmt.Printf("%+v\n", req)

		fmt.Println("=======")
	}
	return p
}
