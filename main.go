package main

import (
  "log"

  "os"
  "os/signal"
  "strconv"

  "github.com/yosssi/gmq/mqtt"
  MqttClient "github.com/yosssi/gmq/mqtt/client"
  InfluxClient "github.com/influxdata/influxdb/client/v2"
)

func handler(influxClient InfluxClient.Client, temp float64) {
  err := writePoint(influxClient, temp)
  if err != nil {
    panic(err)
  }
}

func main() {
  // Set up channel on which to send signal notifications.
  sigc := make(chan os.Signal, 1)
  signal.Notify(sigc, os.Interrupt, os.Kill)

  // Influx
  influxClient, err := getInfluxClient()
  if err != nil {
    log.Fatalln("Error : ", err)
  }
  // MQTT
  mqttClient, err := getMQTTClient()
  if err != nil {
    log.Fatalln("Error : ", err)
  }
  defer mqttClient.Terminate()

  // Subscribe to topics.
  err = mqttClient.Subscribe(&MqttClient.SubscribeOptions{
    SubReqs: []*MqttClient.SubReq{
      &MqttClient.SubReq{
        TopicFilter:  []byte("temperature"),
        QoS:          mqtt.QoS0,
        Handler: func (topic, message []byte) {
          floatString := string(message)[:len(message) - 4]
          i, err := strconv.ParseFloat(string(floatString), 64)
          if err != nil {
            panic(err)
          }
          handler(influxClient, i)
        },
      },
    },
  })
  if err != nil {
    log.Fatalln("Error : ", err)
  }


  // Wait for receiving a signal.
  <-sigc

  // Disconnect the Network Connection.
  if err := mqttClient.Disconnect(); err != nil {
    panic(err)
  }
}
