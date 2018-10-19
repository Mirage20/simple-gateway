/*
 * Copyright (c) 2018 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package gateway

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

type Gateway struct {
	Port       int
	ProxyStore ProxyStore
}

type transport struct {
}

type ProxyStore struct {
	ExactMatch  map[string]*reverseProxy
	PrefixMatch map[string]*reverseProxy
	RegexMatch  map[string]*reverseProxy
}

type Route struct {
	Match       Match
	Destination Destination
}

type Match struct {
	Exact  []string
	Prefix []string
	Regex  []string
}

type Destination struct {
	Host string
	Port int
}

func New(port int, routes []Route) *Gateway {

	//proxies := make(map[string]*reverseProxy)
	gw := &Gateway{
		Port: port,
	}
	gw.UpdateRoutes(routes)
	return gw
}

func (gw *Gateway) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("%+v\n", req.URL.Path)
	proxy := gw.findProxy(req.URL.Path)

	if proxy == nil {
		http.Error(rw, "no route found", http.StatusNotFound)
		return
	}

	fmt.Printf("%p\n", proxy)
	proxy.ServeHTTP(rw, req)
}

func (gw *Gateway) findProxy(path string) *reverseProxy {

	if proxy, ok := gw.ProxyStore.ExactMatch[path]; ok {
		return proxy
	}

	for prefix, proxy := range gw.ProxyStore.PrefixMatch {
		if strings.HasPrefix(path, prefix) {
			return proxy
		}
	}
	return nil
}

func (gw *Gateway) Start() {
	log.Printf("Starting gateway on port %d\n", gw.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", gw.Port), gw); err != nil {
		log.Fatal(err)
	}
}

func (gw *Gateway) UpdateRoutes(routes []Route) {

	proxyStore := ProxyStore{
		ExactMatch:  make(map[string]*reverseProxy),
		PrefixMatch: make(map[string]*reverseProxy),
	}
	for _, route := range routes {
		p := NewReverseProxy(route.Destination)
		for _, path := range route.Match.Exact {
			proxyStore.ExactMatch[path] = p
		}
		for _, path := range route.Match.Prefix {
			proxyStore.PrefixMatch[path] = p
		}
	}
	gw.ProxyStore = proxyStore
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	resp, err := http.DefaultTransport.RoundTrip(r)
	return resp, err
}
