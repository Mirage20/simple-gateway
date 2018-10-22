package filters

import (
	"net/http"
)

type Filter func(http.Handler) http.Handler


func Chain(h http.Handler, filters ...Filter) http.Handler {
	for _, filter := range filters {
		h = filter(h)
	}
	return h
}
