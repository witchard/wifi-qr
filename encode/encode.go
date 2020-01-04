// Package encode can be used to generate the QR code. More details can be found here: https://github.com/zxing/zxing/wiki/Barcode-Contents#wi-fi-network-config-android-ios-11.
package encode

import (
	"fmt"
	"regexp"

	"github.com/skip2/go-qrcode"
)

var escape = regexp.MustCompile(`(["\\;,:])`)

// Escape will correctly escape the string for encoding correctly in the QR code.
func Escape(in string) string {
	return escape.ReplaceAllString(in, `\$1`)
}

// EncodeWpa will return the bytes of a png image encoding the Wi-Fi details
// for the given ssid and password values.
func EncodeWpa(ssid, password string) []byte {
	contents := fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", Escape(ssid), Escape(password))
	png, err := qrcode.Encode(contents, qrcode.Medium, 256)
	if err != nil {
		return nil
	}
	return png
}
