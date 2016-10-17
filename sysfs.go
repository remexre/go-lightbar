package lightbar

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type lightbarImpl struct {
	closed bool
	dir    string
}

func newLightbarImpl(dir string) (*lightbarImpl, error) {
	if info, err := os.Stat(dir); err != nil {
		return nil, err
	} else if !info.IsDir() {
		return nil, errors.New("go-lightbar: not a directory: " + dir)
	}
	return &lightbarImpl{
		closed: false,
		dir:    dir,
	}, nil
}

func (l *lightbarImpl) read(prop string) (string, error) {
	b, err := ioutil.ReadFile(path.Join(l.dir, prop))
	return string(b), err
}
func (l *lightbarImpl) readNums(prop string) ([]byte, error) {
	str, err := l.read(prop)
	if err != nil {
		return nil, err
	}
	numstrs := strings.Split(str, " ")
	nums := make([]byte, len(numstrs))
	for i, numstr := range numstrs {
		n, err := strconv.ParseUint(numstr, 10, 8)
		if err != nil {
			return nil, err
		}
		nums[i] = byte(n)
	}
	return nums, nil
}
func (l *lightbarImpl) write(prop, value string) error {
	if err := ioutil.WriteFile(path.Join(l.dir, "sequence"), []byte("STOP"), 0600); err != nil {
		return err
	}
	return ioutil.WriteFile(path.Join(l.dir, prop), []byte(value), 0600)
}
func (l *lightbarImpl) writeNums(prop string, nums []byte) error {
	numstrs := make([]string, len(nums))
	for i, num := range nums {
		numstrs[i] = strconv.Itoa(int(num))
	}
	return l.write(prop, strings.Join(numstrs, " "))
}

func (l *lightbarImpl) SetBrightness(b byte) error {
	return l.writeNums("brightness", []byte{b})
}

func (l *lightbarImpl) SetLED(led, r, g, b byte) error {
	return l.writeNums("led_rgb", []byte{led, r, g, b})
}

func (l *lightbarImpl) SetLEDs(c [4][3]byte) error {
	return l.writeNums("led_rgb", []byte{
		0, c[0][0], c[0][1], c[0][2],
		1, c[1][0], c[1][1], c[1][2],
		2, c[2][0], c[2][1], c[2][2],
		3, c[3][0], c[3][1], c[3][2],
	})
}

func (l *lightbarImpl) Version() (byte, error) {
	bs, err := l.readNums("version")
	if err != nil {
		return 0, err
	}
	return bs[0], nil
}

func (l *lightbarImpl) FeatureFlags() (byte, error) {
	bs, err := l.readNums("version")
	if err != nil {
		return 0, err
	}
	return bs[1], nil
}
