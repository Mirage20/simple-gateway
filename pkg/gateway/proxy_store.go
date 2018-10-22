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
	"github.com/mirage20/simple-gateway/pkg/gateway/config"
	"github.com/mirage20/simple-gateway/pkg/gateway/proxy"
	"strings"
)

type ProxyStore struct {
	ExactMatch  map[string]*proxy.ReverseProxy
	PrefixMatch map[string]*proxy.ReverseProxy
	RegexMatch  map[string]*proxy.ReverseProxy
}

func (ps *ProxyStore) FindProxy(path string) *proxy.ReverseProxy {
	if p, ok := ps.ExactMatch[path]; ok {
		return p
	}

	for prefix, p := range ps.PrefixMatch {
		if strings.HasPrefix(path, prefix) {
			return p
		}
	}
	return nil
}

func NewProxyStore(routes []config.Route) *ProxyStore {
	ps := &ProxyStore{
		ExactMatch:  make(map[string]*proxy.ReverseProxy),
		PrefixMatch: make(map[string]*proxy.ReverseProxy),
		RegexMatch:  make(map[string]*proxy.ReverseProxy),
	}

	for _, route := range routes {
		p := proxy.NewReverseProxy(route.Destination.Host,route.Destination.Port)
		for _, path := range route.Match.Exact {
			ps.ExactMatch[path] = p
		}
		for _, path := range route.Match.Prefix {
			ps.PrefixMatch[path] = p
		}
	}
	return ps
}
