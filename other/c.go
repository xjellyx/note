package main

import (
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/getFile", demo)
	http.ListenAndServe(":8080", nil)

}

func demo(w http.ResponseWriter, r *http.Request) {
	file := "/home/srlemon/Downloads/NVIDIA-Linux-x86_64-430.40.run"
	_, err := os.Stat(file)
	if ok := os.IsExist(err); !ok {
		http.NotFound(w, r)
	}
	http.ServeFile(w, r, file)
}
