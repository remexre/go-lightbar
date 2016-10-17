package lightbar

// Lightbar is an interface to represent the system lightbar.
type Lightbar interface {
	SetBrightness(b byte) error
	SetLED(led, r, g, b byte) error
	SetLEDs(colors [4][3]byte) error

	Version() (byte, error)
	FeatureFlags() (byte, error)
}

var lightbar Lightbar

func init() {
	var err error
	lightbar, err = newLightbarImpl("/sys/class/chromeos/cros_ec/lightbar")
	if err != nil {
		panic(err)
	}
}

// Get returns the Lightbar for this system.
func Get() Lightbar { return lightbar }
