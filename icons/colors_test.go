package icons

import "testing"

func TestHexToRGB(t *testing.T) {
	testcases := []struct {
		hex string
		rgb [3]uint8
		err error
	}{
		{"#000000", [3]uint8{0, 0, 0}, nil},
		{"#ffffff", [3]uint8{255, 255, 255}, nil},
		{"#ff0000", [3]uint8{255, 0, 0}, nil},
		{"#00ff00", [3]uint8{0, 255, 0}, nil},
		{"#0000ff", [3]uint8{0, 0, 255}, nil},
		{"#994943", [3]uint8{153, 73, 67}, nil},
		{"#000", [3]uint8{0, 0, 0}, nil},
		{"#fff", [3]uint8{255, 255, 255}, nil},
		{"#f00", [3]uint8{255, 0, 0}, nil},
		{"#0f0", [3]uint8{0, 255, 0}, nil},
		{"#00f", [3]uint8{0, 0, 255}, nil},
		{"#0000000", [3]uint8{0, 0, 0}, errInvalidFormat},
		{"#fffffff", [3]uint8{0, 0, 0}, errInvalidFormat},
	}
	for _, tc := range testcases {
		t.Run(tc.hex, func(t *testing.T) {
			rgb, err := hexToRGB(tc.hex)
			if err != tc.err {
				t.Errorf("got %v; want %v", err, tc.err)
			}
			if rgb != tc.rgb {
				t.Errorf("got %v; want %v", rgb, tc.rgb)
			}
		})
	}
}
