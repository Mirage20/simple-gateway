package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type ReverseProxy struct {
	httputil.ReverseProxy
	Host string
	Port int
}

func NewReverseProxy(host string, port int) *ReverseProxy {
	p := &ReverseProxy{
		Host: host,
		Port: port,
		ReverseProxy: httputil.ReverseProxy{
			Director: director(host, port),
		},
	}
	//p.Director = director(p.Host, p.Port)
	//p.Director = director(host, port)
	return p
}

func director(host string, port int) func(req *http.Request) {
	return func(req *http.Request) {
		req.URL.Scheme = "http"
		host := fmt.Sprintf("%s:%d", host, port)
		req.URL.Host = host
		req.Host = host

		//b, err := httputil.DumpRequest(req, true)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Println(string(b))
		//fmt.Printf("%+v\n", req)
		//
		//fmt.Println("=======")
	}
}
