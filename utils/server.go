package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/log"
)

func GetHost(host string) string {
	if host == "0" || host == "" {
		return ""
	}
	if _, err := strconv.ParseInt(host, 10, 64); err == nil {
		host = fmt.Sprintf(":%s", host)
	}
	return host
}

// host = 0.0.0.0:8080
// host = 8080
// host =0 return
func ServerFromMux(mux http.Handler, host string) {
	host = GetHost(host)
	if host == "" {
		return
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

func HTTPServerWriteErr(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(ToJsonBytes(map[string]string{"message": err.Error()}))
}
func HTTPServerWriteSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(ToJsonBytes(map[string]interface{}{"data": data}))
}

func HTTPServerCORSMiddleware(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS,PUT")
		// w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		fn(w, r)
	}
}
