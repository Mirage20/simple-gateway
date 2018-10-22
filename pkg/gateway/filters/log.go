package filters

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func Log() Filter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := httputil.DumpRequest(r, true)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(b))
			defer func() {
				fmt.Printf("%+v\n", w)
			}()
			h.ServeHTTP(w, r)
		})
	}
}
