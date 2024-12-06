package gpio

import (
    "testing"
    "time"
    "fmt"
)

func TestDigitalOutput(t *testing.T) {
    defer Close()

    do := NewDigitalOut(18, Low)
    defer do.Close()

    exeTime := 5.0
    dt := 0.01
    ticker := time.NewTicker(time.Duration(dt * float64(time.Second)))
    defer ticker.Stop()

    toggled := false

    for {
        <- ticker.C
        
        if exeTime <= 0 {
            break
        }

        if exeTime <= 2.5 && !toggled {
            do.Toggle()
            toggled = true
        }
        
        exeTime -= dt
        fmt.Printf("\r%.2f, %v", exeTime, do.Read())
    }
}

func TestPwm(t *testing.T) {
    defer Close()

    pwm, err := NewPwm(18)
    if err != nil {
        t.Fatal(err)
    }
    defer pwm.Close()

    exeTime := 5.0
    dt := 0.01
    ticker := time.NewTicker(time.Duration(dt * float64(time.Second)))
    defer ticker.Stop()

    for {
        <- ticker.C
        
        if exeTime <= 0 {
            break
        }

        dutyCycle := exeTime / 5.0 * 100 + 1
        pwm.Write(int(dutyCycle))

        exeTime -= dt
        fmt.Printf("\r%.2f, %.2f", exeTime, dutyCycle)
    }
}
