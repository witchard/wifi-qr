package main

import (
	"bytes"
	"encoding/base64"
	"syscall/js"

	"github.com/witchard/wifi-qr/encode"
)

func main() {
	ssid := js.Global().Get("document").Call("getElementById", "ssid").Get("value").String()
	pass := js.Global().Get("document").Call("getElementById", "password").Get("value").String()
	qr := encode.EncodeWpa(ssid, pass)
	buf := bytes.Buffer{}
	enc := base64.NewEncoder(base64.StdEncoding, &buf)
	enc.Write(qr)
	enc.Close()
	img := "data:image/png;base64," + buf.String()
	js.Global().Get("document").Call("getElementById", "out").Set("src", img)
}
