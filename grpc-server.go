package main

import (
    "google.golang.org/grpc"
    "log"
    "net"
    service "./service"
)

func main() {
    listenPort, err := net.Listen("tcp", ":2525")
    if err != nil {
        log.Fatalln(err)
    }
    server := grpc.NewServer()
    echoService := &service.EchoService{}
    service.RegisterEchoServer(server, echoService)
    err = server.Serve(listenPort)
    if err != nil {
        log.Fatalln(err)
    }
}

