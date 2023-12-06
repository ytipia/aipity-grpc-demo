package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	// import proto
	pb_group "aipity/proto/group"
	pb_user "aipity/proto/user"

	// import service
	"aipity/service/group"
	"aipity/service/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port    = flag.Int("port", 50051, "The server port")
	crtFile = filepath.Join("certs", "server.crt")
	keyFile = filepath.Join("certs", "server.key")
	caFile  = filepath.Join("certs", "ca.crt")
)

func init() {

	// config log to file and console both
	file, _ := os.OpenFile("aipity-grpc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(io.MultiWriter(file, os.Stdout))

}

func main() {
	flag.Parse()

	// mTLS settings
	certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append client certs")
	}

	opts := []grpc.ServerOption{
		// Enable TLS for all incoming connections.
		grpc.Creds( // Create the TLS credentials
			credentials.NewTLS(&tls.Config{
				ClientAuth:   tls.RequireAndVerifyClientCert,
				Certificates: []tls.Certificate{certificate},
				ClientCAs:    certPool,
			},
			)),
	}

	s := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb_user.RegisterUserServiceServer(s, &user.Server{})
	pb_group.RegisterGroupServiceServer(s, &group.Server{})
	log.Printf("grpc server with mTLS enabled listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
