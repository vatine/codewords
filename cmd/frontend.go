// Command to start codewords frontend
package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	
	"github.com/vatine/codewords/frontend"

	//	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
 	"golang.org/x/net/context"
)

var port = flag.String("-port", ":8080", "Port to listen to")
var backend = flag.String("-backend", "localhost:8090", "Backend URL")

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><head><title>Codewords</title></head>")
	fmt.Fprintln(w, "<body><h1>Welcome to the Codewords service</h1>")
	renderForm(w, r)
	fmt.Fprintln(w, "</body></html>")
}

func renderForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<form action="/getcode"><input type="submit" value="Get code word"></form>`)
}

func getCodeword(w http.ResponseWriter, r *http.Request) {
	codeword, err := frontend.NextCodeword(context.Background())
	if err != nil {
		fmt.Printf("Error getting codeword, %s", err)
	}
	fmt.Fprintln(w, "<html><head><title>Codewords</title></head>")
	fmt.Fprintln(w, "<body><h1>Welcome to the Codewords service</h1>")
	fmt.Fprintf(w, "Your codeword is <b>%s</b><p>\n", codeword)
	renderForm(w, r)
	fmt.Fprintln(w, "</body></html>")
	
}

func main () {
	flag.Parse()
	
	_, err := frontend.Connect(*backend, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Error connecting to the backend, %s\n", err)
		return
	}
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/getcode", getCodeword)
	err = http.ListenAndServe(*port, nil)
	if err != nil {
		fmt.Println("Error starting listener, %s", err)
	}
}
