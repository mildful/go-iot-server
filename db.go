package main

import (
  "fmt"
  "time"

  "github.com/influxdata/influxdb/client/v2"
)

const (
  DBName = "iot"
  username = "root"
  password = "root"
)

func getInfluxClient() (client.Client, error) {
  c, err := client.NewHTTPClient(client.HTTPConfig{
    Addr: "http://localhost:8086",
    Username: username,
    Password: password,
  })
  if err != nil {
    return nil, err
  }

  return c, nil
}

func writePoint(c client.Client, temp float64) (error) {
  // Create a new point batch.
  bp, err := client.NewBatchPoints(client.BatchPointsConfig{
    Database: DBName,
    Precision: "s",
  })
  if err != nil {
    return err
  }

  // Create a point and add to batch.
  tags := map[string]string{"displayName": "temperature"}
  fields := map[string]interface{}{
    "value": temp,
  }
  pt, err := client.NewPoint("temperature", tags, fields, time.Now())
  if err != nil {
    return err
  }

  bp.AddPoint(pt)
  err = c.Write(bp)
  if err != nil {
    return err
  }
  return nil
}
