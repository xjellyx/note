package main

import (
	"github.com/nfnt/resize"
	"image"
	"image/png"
	_ "image/png"
	"os"
)

//
func main() {
	ReSizeImgSave("/home/olongfen/data/file67/analysis_information_data/result/GF02_PA1_040680_20220227_KS450_01_018_L1A_01/GF02_PA1_040680_20220227_KS450_01_018_L1A_01.png",
		"/home/olongfen/data/file67/analysis_information_data/result/GF02_PA1_040680_20220227_KS450_01_018_L1A_01/GF02_PA1_040680_20220227_KS450_01_018_L1A_01.png",
		2048, 2048)
}

func ReSizeImgSave(in, out string, w, h uint) (err error) {
	var (
		file    *os.File
		img     image.Image
		outFile *os.File
	)
	file, err = os.Open(in)
	if err != nil {
		return
	}

	// decode jpeg into image.Image
	img, err = png.Decode(file)
	if err != nil {
		return
	}
	file.Close()

	m := resize.Resize(w, h, img, resize.Lanczos3)
	if outFile, err = os.Create(out); err != nil {
		return
	}
	if err = png.Encode(outFile, m); err != nil {
		return
	}

	return

}
