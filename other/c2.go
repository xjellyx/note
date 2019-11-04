package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func gethash(path string) (hash string) {
	file, err := os.Open(path)
	if err == nil {
		h_ob := sha1.New()
		_, err = io.Copy(h_ob, file)
		if err == nil {
			hash := h_ob.Sum(nil)
			hashvalue := hex.EncodeToString(hash)
			_=os.Rename(path,hashvalue)
			return hashvalue
		}
	}

	defer file.Close()
	return
}
func main() {
	path := "/data/fedora-data/gocode/src/github.com/srlemon/note/1.jpeg"
	//path:="md5.go"
	hash := gethash(path)
	fmt.Printf("%s hash: %s \n", path, hash)
}