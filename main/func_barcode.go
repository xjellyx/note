package main

import (
	"bytes"
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Create the barcode
	qrCode, _ := qr.Encode("http://www.olongfen.ltd:80", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	// create the output file
	buf := new(bytes.Buffer)
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(buf, qrCode)
	fmt.Println(string(buf.Bytes()))

}
