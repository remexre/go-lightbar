package lightbarpc

import lightbar "github.com/remexre/go-lightbar"

// LightbarRPC is a struct to allow RPCs for the lightbar using the net/rpc
// library.
type LightbarRPC struct {
	Lightbar lightbar.Lightbar
}

// Get returns the Lightbar for this system, wrapped in a LightbarRPC.
func Get() *LightbarRPC {
	return &LightbarRPC{lightbar.Get()}
}

// SetBrightness sets the brightness.
func (l *LightbarRPC) SetBrightness(brightness byte, _ *struct{}) error {
	return l.Lightbar.SetBrightness(brightness)
}

// SetLED sets the colors of an LED.
func (l *LightbarRPC) SetLED(info [4]byte, _ *struct{}) error {
	return l.Lightbar.SetLED(info[0], info[1], info[2], info[3])
}

// SetLEDs sets the colors of the LEDs.
func (l *LightbarRPC) SetLEDs(colors [4][3]byte, _ *struct{}) error {
	return l.Lightbar.SetLEDs(colors)
}
