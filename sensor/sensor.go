package sensor

type Sensor interface {
    Read() float64
    Close()
}
