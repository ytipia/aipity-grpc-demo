package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	// import proto
	pb_group "aipity/proto/group"
	pb_user "aipity/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	address  = "localhost:50051"
	hostname = "localhost"
	crtFile  = filepath.Join("certs", "client.crt")
	keyFile  = filepath.Join("certs", "client.key")
	caFile   = filepath.Join("certs", "ca.crt")
)

func main() {

	// Load the client certificates from disk
	certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(caFile)
	if err != nil {
		log.Fatalf("could not read ca certificate: %s", err)
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("failed to append ca certs")
	}

	// log.SetFlags
	file, _ := os.OpenFile("aipity-grpc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(io.MultiWriter(file, os.Stdout))

	opts := []grpc.DialOption{
		// transport credentials.
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:   hostname, // NOTE: this is required!
			Certificates: []tls.Certificate{certificate},
			RootCAs:      certPool,
		})),
	}

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect to grpc server: %v", err)
	}
	defer conn.Close()
	c := pb_user.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	// 1. createuser
	name := "aipity"
	password := "123456"
	email := "whsasf@aipity.com"
	phone := "111111111"
	var status int32 = 1
	var role int32 = 1
	createTime := time.Now().Unix()
	var theme int32 = 1
	var language int32 = 1

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateUser(ctx, &pb_user.User{
		Name:       name,
		Password:   password,
		Email:      email,
		Phone:      phone,
		Status:     status,
		CreateTime: createTime,
		Theme:      theme,
		Role:       role,
		Language:   language,
	})
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}
	log.Printf("creatd user: %s added successfully", r)

	// 2. create group
	g := pb_group.NewGroupServiceClient(conn)
	gname := "aipity"
	gid := 1

	t, err := g.CreateGroup(ctx, &pb_group.Group{
		Name: gname,
		Id:   int32(gid),
	})
	if err != nil {
		log.Fatalf("Could not create group: %v", err)
	}
	log.Printf("created group: %s added successfully", t)
}
