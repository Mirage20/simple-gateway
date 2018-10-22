package filters

import (
	"fmt"
	"net/http"
	"time"
)

func Time() Filter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				fmt.Printf("%q request took %v\n", r.URL.RequestURI(), time.Since(start))
			}()
			h.ServeHTTP(w, r)
		})
	}
}
