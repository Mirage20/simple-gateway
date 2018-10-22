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
	"github.com/mirage20/simple-gateway/pkg/gateway/config"
	"github.com/mirage20/simple-gateway/pkg/gateway/filters"
	"log"
	"os"

	//"github.com/golang/glog"

)

const (
	threadsPerController = 2
)

var (
	port        int
	routeConfig string
)

func main() {
	flag.Parse()

	if routeConfig == "" {
		log.Fatal("Please specify a route config using -routes flag")
		os.Exit(1)
	}
	//stopCh := signals.SetupSignalHandler()

	r, err := config.Load(routeConfig)
	if err != nil {
		log.Fatal("Error while loading route config: ", err)
		os.Exit(1)
	}
	//r := []config.Route{
	//	{
	//		Match: config.Match{
	//			Exact:  []string{"/foo", "/aaa"},
	//			Prefix: []string{"/products"},
	//		},
	//		Destination: config.Destination{
	//			Host: "google.lk",
	//			Port: 80,
	//		},
	//	},
	//	{
	//		Match: config.Match{
	//			Exact:  []string{"/bar", "/bbb"},
	//			Prefix: []string{"/users"},
	//		},
	//		Destination: config.Destination{
	//			Host: "example.com",
	//			Port: 80,
	//		},
	//	},
	//}

	gw := gateway.NewGateway(port, r, filters.Log(), filters.Time())

	gw.Start()

	// Prevent exiting the main process
	//<-stopCh
}

func init() {
	flag.IntVar(&port, "port", 8080, "Gateway listening port")
	flag.StringVar(&routeConfig, "routes", "", "Gateway route config file")
}
