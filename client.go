package main
import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "log"
    service "./service"
)

func main() {
    //sampleなのでwithInsecure
    conn, err := grpc.Dial("127.0.0.1:2525", grpc.WithInsecure())
    if err != nil {
        log.Fatal("client connection error:", err)
    }
    defer conn.Close()
    client := service.NewEchoClient(conn)
    message := &service.GetEchoMessage{TargetEcho: "hiraki"}
    res, err := client.GetEcho(context.TODO(), message)
    fmt.Printf("result:%#v \n", res)
    fmt.Printf("error::%#v \n", err)
}

