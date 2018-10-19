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

package main

import (
	"flag"
	"github.com/mirage20/simple-gateway/pkg/gateway"
	//"github.com/golang/glog"

)

const (
	threadsPerController = 2
)

var (
	port       int
	kubeconfig string
)

func main() {
	flag.Parse()

	//stopCh := signals.SetupSignalHandler()

	r := []gateway.Route{
		{
			Match: gateway.Match{
				Exact: []string{"/foo", "/aaa"},
				Prefix:[]string{"/products"},
			},
			Destination: gateway.Destination{
				Host: "google.lk",
				Port: 80,
			},
		},
		{
			Match: gateway.Match{
				Exact: []string{"/bar", "/bbb"},
				Prefix:[]string{"/users"},
			},
			Destination: gateway.Destination{
				Host: "example.com",
				Port: 80,
			},
		},
	}

	gw := gateway.New(port, r)

	gw.Start()

	// Prevent exiting the main process
	//<-stopCh
}

func init() {
	flag.IntVar(&port, "port", 8080, "Gateway listening port")
}
