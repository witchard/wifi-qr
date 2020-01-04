# WiFi-QR

A simple utility for generating QR codes for your wifi login.

Based on details from: https://feeding.cloud.geek.nz/posts/encoding-wifi-access-point-passwords-qr-code/.

## Command line tool

Install: `go get -u github.com/witchard/wifi-qr`.

Usage: `wifi-qr -s <SSID> -p <Password> > qr.png`

## Go package

```go
import "github.com/witchard/wifi-qr/encode

...
	png_bytes := encode.EncodeWpa("ssid", "password")
...
```

## Wasm

In the `wasm` folder you will find a simple web browser example using wasm. Go into that folder and run `make serve` to see it in action.
