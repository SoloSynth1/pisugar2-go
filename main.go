package main

import (
	"fmt"
	i2c "github.com/d2r2/go-i2c"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Battery struct {
	high uint8
	low  uint8
}

type BatteryCurvePoint struct {
	voltage float32
	level   float32
}

func getBatteryCurve() []BatteryCurvePoint {
	return []BatteryCurvePoint{
		{4.10, 100.0},
		{4.05, 95.0},
		{3.90, 88.0},
		{3.80, 77.0},
		{3.70, 65.0},
		{3.62, 55.0},
		{3.58, 49.0},
		{3.49, 25.6},
		{3.32, 4.5},
		{3.1, 0.0},
	}
}

var batteryCurve = getBatteryCurve()

func (b *Battery) GetVoltage() float32 {
	var v float32
	if b.high&0x20 != 0 {
		b.low = ^b.low & 0xff
		b.high = ^b.high & 0x1f
		v = float32(uint16(b.high|0b1100_0000)<<8) + float32(b.low)
		v = (2600.0 - v*0.26855) / 1000.
	} else {
		v = float32(uint16(b.high&0x1f)<<8) + float32(b.low)
		v = (2600.0 + v*0.26855) / 1000.
	}
	return v
}

func (b *Battery) GetPercentage() float32 {
	v := b.GetVoltage()
	for i, point := range batteryCurve {
		if v >= point.voltage {
			if i == 0 {
				return point.level
			} else {
				vHigh := batteryCurve[i-1].voltage
				lHigh := batteryCurve[i-1].level
				percent := (v - point.voltage) / (vHigh - point.voltage)
				return point.level + percent*(lHigh-point.level)
			}
		}
	}
	return 0.0
}

func main() {
	i2cBus := 1
	i2c, err := i2c.NewI2C(0x75, i2cBus)
	check(err)
	defer i2c.Close()
	high, err := i2c.ReadRegU8(0xd1)
	check(err)
	low, err := i2c.ReadRegU8(0xd0)
	check(err)
	b := Battery{high, low}
	fmt.Printf("voltage is: %v V\n", b.GetVoltage())
	fmt.Printf("Percentage: %v %\n", b.GetPercentage())
}
