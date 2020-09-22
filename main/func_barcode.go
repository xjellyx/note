package main

import (
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	// Create the barcode
	qrCode, _ := qr.Encode("http://b319.photo.store.qq.com/psb?/V11OLDdE1l74ez/Rn9P8zCw*TvCUL*nO.QyLMmZVhnzKZagf7Pqt3wjZHI!/m/dD8BAAAAAAAAnull&bo=KgM4BCoDOAQRBzA!&rf=photolist&t=5", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

}
