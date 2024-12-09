package gpio

import (
	"fmt"
	"os"

    "github.com/vistormu/goraspio/num"
)

 // #include <linux/i2c-dev.h>
import "C"

const (
	I2C_SLAVE = C.I2C_SLAVE
)

type I2C struct {
	addr uint8
	bus  int
	rc   *os.File
}

func NewI2C(addr uint8, bus int) (*I2C, error) {
	f, err := os.OpenFile(fmt.Sprintf("/dev/i2c-%d", bus), os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	if err := ioctl(f.Fd(), I2C_SLAVE, uintptr(addr)); err != nil {
		return nil, err
	}
	v := &I2C{rc: f, bus: bus, addr: addr}
	return v, nil
}

func (i *I2C) Read(registers []byte, nBytes []int) ([]byte, error) {
    result := make([]byte, num.Sum(nBytes))
    for j := range len(registers) {
        reg := registers[j]
        n := nBytes[j]

        _, err := i.rc.Write([]byte{reg})
        if err != nil {
            return nil, err
        }

        buf := make([]byte, n)
        _, err = i.rc.Read(buf)
        if err != nil {
            return nil, err
        }

        for i, value := range buf {
            result[j+i] = value
        }
    }

    return result, nil
}

func (i *I2C) Write(reg byte, value byte) error {
	buf := []byte{reg, value}
	_, err := i.rc.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

func (i *I2C) Close() error {
	return i.rc.Close()
}