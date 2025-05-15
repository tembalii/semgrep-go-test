package insecuregrpc

import (
    "google.golang.org/grpc"
		"google.golang.org/grpc/credentials"
		"google.golang.org/grpc/credentials/insecure"
)

func unsafe() {
    // ruleid: grpc-client-new-insecure-connection
    conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
}

func safe() {
    // ok: grpc-client-new-insecure-connection
    conn, err := grpc.Dial(address, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
}
