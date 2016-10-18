package main

import (
	"fmt"
	"math"
	"os"
	"time"

	lightbar "github.com/remexre/go-lightbar"
)

func main() {
	lb := lightbar.Get()
	for now := range time.Tick(50 * time.Millisecond) {
		hue := int(now.UnixNano()/2e7) % 360
		colors := [4][3]byte{
			hueToRGB(hue),
			hueToRGB(hue + 90),
			hueToRGB(hue + 180),
			hueToRGB(hue + 270),
		}
		if err := lb.SetLEDs(colors); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func hueToRGB(hue int) [3]byte {
	hue %= 360
	h := float64(hue) / 60.0
	hmod2 := h
	for hmod2 > 2.0 {
		hmod2 -= 2.0
	}
	x := 1.0 - math.Abs(hmod2-1)

	if h < 1 {
		return [3]byte{255, byte(255 * x), 0}
	} else if h < 2 {
		return [3]byte{byte(255 * x), 255, 0}
	} else if h < 3 {
		return [3]byte{0, 255, byte(255 * x)}
	} else if h < 4 {
		return [3]byte{0, byte(255 * x), 255}
	} else if h < 5 {
		return [3]byte{byte(255 * x), 0, 255}
	}
	// h < 6
	return [3]byte{255, 0, byte(255 * x)}
}
