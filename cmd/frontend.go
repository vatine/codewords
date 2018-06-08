// Command to start codewords frontend
package main

import (
	"flag"
	"fmt"
	"net/http"
	
	"github.com/vatine/codewords/frontend"

	//	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
 	"golang.org/x/net/context"
)

var port = flag.String("-port", ":8080", "Port to listen to")
var backend = flag.String("-backend", "http://localhost:8081", "Backend URL")

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><head><title>Codewords</title></head>")
	fmt.Fprintln(w, "<body><h1>Welcome to the Codewords service</h1>")
	renderForm(w, r)
	fmt.Fprintln(w, "</body></html>")
}

func renderForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<form action=...")
}

func getCodeword(w http.ResponseWriter, r *http.Request) {
	codeword, _ := frontend.NextCodeword(context.Background())
	fmt.Fprintln(w, "<html><head><title>Codewords</title></head>")
	fmt.Fprintln(w, "<body><h1>Welcome to the Codewords service</h1>")
	fmt.Fprintf(w, "Your codeword is <b>%s</b><p>\n", codeword)
	renderForm(w, r)
	fmt.Fprintln(w, "</body></html>")
	
}

func main () {
	flag.Parse()
	
	frontend.Connect(*backend)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(*port, nil)
}
