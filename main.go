package main

import (
  "fmt"
  "os"
  "os/signal"

  "github.com/yosssi/gmq/mqtt"
  "github.com/yosssi/gmq/mqtt/client"
)

func handler(topic, message []byte) {
  fmt.Println(string(message))
}

func main() {
  // Set up channel on which to send signal notifications.
  sigc := make(chan os.Signal, 1)
  signal.Notify(sigc, os.Interrupt, os.Kill)

  // Create an MQTT Client.
  cli := client.New(&client.Options{
    // Define the processing of the error handler.
    ErrorHandler: func(err error) {
      fmt.Println(err)
    },
  })

  // Terminate the cliebt.
  defer cli.Terminate()

  // Connect toe the MQTT server.
  err := cli.Connect(&client.ConnectOptions{
    Network: "tcp",
    Address: "0.0.0.0:1883",
    ClientID: []byte("example"),
  })
  if err != nil {
    panic(err)
  }

  // Subscribe to topics.
  err = cli.Subscribe(&client.SubscribeOptions{
    SubReqs: []*client.SubReq{
      &client.SubReq{
        TopicFilter:  []byte("temp"),
        QoS:          mqtt.QoS0,
        Handler:      handler,
      },
    },
  })
  if err != nil {
    panic(err)
  }
  /*err = cli.Publish(&client.PublishOptions{
        QoS:       mqtt.QoS0,
        TopicName: []byte("temp"),
        Message:   []byte("11"),
    })
    if err != nil {
        panic(err)
    }*/

  // Unsubscribe from topics.
  /*err = cli.Unsubscribe(&client.UnsubscribeOptions{
    TopicFilters: [][]byte{
      []byte("temp"),
    },
  })
  if err != nil {
    panic(err)
  }*/

  // Wait for receiving a signal.
  <-sigc

  // Disconnect the Network Connection.
  if err := cli.Disconnect(); err != nil {
    panic(err)
  }
}
