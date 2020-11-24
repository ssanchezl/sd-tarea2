package main

import (
  "log"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/dcordova/sd_tarea2/data_service"

  //"strings"
  //"bufio"
  //"encoding/csv"
  //"fmt"
  //"io"
  //"os"
  //"strconv" // Conversion de strings a int y viceversa
)


func main() {

  // Conectarse como cliente al servidor //
  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":9000", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("Could not connect: %s", err)
  }
  defer conn.Close()

  c := data_service.NewDataServiceClient(conn)

  // Hello world
  message := data_service.Message{
    Body: "Hello from the client!",
  }

  response, err := c.SayHello(context.Background(), &message)
  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }

  log.Printf("Response from Server: %s", response.Body)
}