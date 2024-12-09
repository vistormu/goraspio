package sensor

import (
    "testing"
    "time"
    "fmt"
)


func TestSensor(t *testing.T) {
    sensorName := "nse5310"
    var sensor Sensor
    var err error
    switch sensorName {
    case "mcp3201":
        sensor, err = NewMcp3201(5.0, 24)
    case "nse5310":
        sensor, err = NewNse5310(0x40, 1)
    default:
        t.Fatal("unknown sensor")
    }
    if err != nil {
        t.Fatal(err)
    }
    defer sensor.Close()

    exeTime := 10.0
    dt := 0.001
    ticker := time.NewTicker(time.Duration(dt*float64(time.Second)))
    defer ticker.Stop()

    for range int(exeTime/dt) {
        <- ticker.C

        value := sensor.Read()

        fmt.Println(value)
    }
}
