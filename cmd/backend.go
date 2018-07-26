// Command to start codewords frontend
package main

import (
	"fmt"
	"flag"
	"net"
	"net/http"
	
	"github.com/vatine/codewords/backend"
	cw "github.com/vatine/codewords"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port = flag.String("-port", ":8090", "Port to listen to, for grpc")
var promPort = flag.String("-monitoring", ":8081", "Port to listen to, for monitoring")

func main () {
	flag.Parse()
	s := grpc.NewServer()
	
	cw.RegisterCodewordsServiceServer(s, &backend.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	listener, err := net.Listen("tcp", *port)
	if err != nil {
		fmt.Printf("Unable to alocate GRPC service port, %s\n", err)
		return
	}
	
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(*promPort, nil)
	if err := s.Serve(listener); err != nil {
		fmt.Printf("Failed to start grpc server, %s\n", err)
	}
}
