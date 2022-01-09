package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	// homework 4 return 200 on healthz
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200 ok\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	w.Header().Set("content-type", "application/json")
	// homework 1: add request header to response header
	for k, v := range r.Header {
		w.Header().Add(k, v[0])
	}
	// homework 2: get version and add to header
	os.Setenv("VERSION", "1")
	w.Header().Add("VERSION", os.Getenv("VERSION"))

	// homework 3: log user ip and response code
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		fmt.Println("user ip is: ", netIP)
	}
	// not sure how to get http response and log
	fmt.Println("successfully served one request with 200 response code")

	user := r.URL.Query().Get("user")
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}

}
