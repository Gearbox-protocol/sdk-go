package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/log"
)

// host = 0.0.0.0:8080
// host = 8080
// host =0 return
func ServerFromMux(mux *http.ServeMux, host string) {
	if host == "0" || host == "" {
		return
	}
	if _, err := strconv.ParseInt(host, 10, 64); err == nil {
		host = fmt.Sprintf(":%s", host)
	}
	srv := http.Server{
		Addr:    host,
		Handler: mux,
	}

	go func() {
		log.Infof("Starting prometheus at %s", host)
		srv.ListenAndServe()
	}()
}
