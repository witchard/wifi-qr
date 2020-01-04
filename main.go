package main

import (
	"flag"
	"github.com/witchard/wifi-qr/encode"
	"log"
	"os"
)

func main() {
	ssid := flag.String("s", "", "Wi-Fi SSID")
	pass := flag.String("p", "", "Wi-Fi Password")
	flag.Parse()
	if len(*ssid) == 0 || len(*pass) == 0 {
		log.Fatalln("-s and -p are both required")
	}
	data := encode.EncodeWpa(*ssid, *pass)
	os.Stdout.Write(data)
}
