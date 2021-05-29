package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	remote, err := url.Parse("https://storage.googleapis.com")
	if err != nil {
		panic(err)
	}

	log.Printf("starting proxy server for GCS volumes on port %s...", port)

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	if err = http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.Header().Set("Access-Control-Allow-Origin", "*") // Handle CORS bug
		p.ServeHTTP(w, r)
	}
}
