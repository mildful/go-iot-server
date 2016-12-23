package main

//"github.com/yosssi/gmq/mqtt"
import (
  "fmt"

  "github.com/yosssi/gmq/mqtt/client"
)

func getMQTTClient() (*client.Client, error) {
  cli := client.New(&client.Options{
    // Define the processing of the error handler.
    ErrorHandler: func(err error) {
      fmt.Println(err)
    },
  })
  // Connect toe the MQTT server.
  err := cli.Connect(&client.ConnectOptions{
    Network: "tcp",
    Address: "0.0.0.0:1883",
    ClientID: []byte("example"),
  })
  if err != nil {
    return nil, err
  }

  return cli, nil
}
