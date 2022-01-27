package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"satya.com/satya-dillikar/grpcexample/interceptors"
	pingpong "satya.com/satya-dillikar/grpcexample/pingpong"
)

func main() {

	// Add PingCounter from the interceptors package
	pc := interceptors.PingCounter{}
	apiserver := grpc.NewServer(
		// Here we add the chaining of interceptors, they will execute in order
		grpc.ChainUnaryInterceptor(
			pc.ServerCount,
			interceptors.VerifyUnaryServer,
			interceptors.LogRequest,
		),
	)

	// Start listening on a TCP Port
	lis, err := net.Listen("tcp", "127.0.0.1:9990")
	if err != nil {
		log.Fatal(err)
	}
	// We need to tell the code WHAT TO do on each request, ie. The business logic.
	// In GRPC cases, the Server is acutally just an Interface
	// So we need a struct which fulfills the server interface
	// see server.go
	s := &Server{}
	// Register the API server as a PingPong Server
	// The register function is a generated piece by protoc.
	pingpong.RegisterPingPongServer(apiserver, s)
	// Start serving in a goroutine to not block
	go func() {
		log.Fatal(apiserver.Serve(lis))
	}()
	// Wrap the GRPC Server in grpc-web and also host the UI
	grpcWebServer := grpcweb.WrapServer(apiserver)
	// Lets put the wrapped grpc server in our multiplexer struct so
	// it can reach the grpc server in its handler
	multiplex := grpcMultiplexer{
		grpcWebServer,
	}

	// We need a http router
	r := http.NewServeMux()
	// Load the static webpage with a http fileserver
	webapp := http.FileServer(http.Dir("ui/pingpongapp/build"))
	// Host the Web Application at /, and wrap it in the GRPC Multiplexer
	// This allows grpc requests to transfer over HTTP1. then be
	// routed by the multiplexer
	r.Handle("/", multiplex.Handler(webapp))
	// Create a HTTP server and bind the router to it, and set wanted address
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8090",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("server started on port :8090")
	// Serve the webapp over TLS
	//log.Fatal(srv.ListenAndServeTLS("cert/server.crt", "cert/server.key"))
	log.Fatal(srv.ListenAndServe())
}

/* // GenerateTLSApi will load TLS certificates and key and create a grpc server with those.
func GenerateTLSApi(pemPath, keyPath string) (*grpc.Server, error) {
	cred, err := credentials.NewServerTLSFromFile(pemPath, keyPath)
	if err != nil {
		return nil, err
	}

	// Add PingCounter from the interceptors package
	//pc := interceptors.PingCounter{}
	s := grpc.NewServer(
		grpc.Creds(cred),
		// Here we add the chaining of interceptors, they will execute in order
		grpc.ChainUnaryInterceptor(
			//pc.ServerCount,
			interceptors.VerifyUnaryServer,
			interceptors.LogRequest,
		),
	)
	return s, nil
} */

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

// Handler is used to route requests to either grpc or to regular http
func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
