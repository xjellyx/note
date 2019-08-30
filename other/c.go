package main

import (
	"fmt"
	// "fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", SayHello3)
	serveMux.HandleFunc("/bye", SayBye)
	// serveMux.HandleFunc("/static", StaticServer)

	server := http.Server{
		Addr:        ":8080",
		Handler:     serveMux,
		ReadTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func SayHello3(w http.ResponseWriter, r *http.Request) {
	if ok, _ := regexp.MatchString("/static/", r.URL.String()); ok {
		StaticServer(w, r)
		return
	}
	io.WriteString(w, "hello world")
}

func SayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Byebye")
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.StripPrefix("/static/",
		http.FileServer(http.Dir(wd))).ServeHTTP(w, r)
	var (
		data []byte
	)
	if data, err = ioutil.ReadAll(r.Body); err != nil {
		return
	}
	defer r.Body.Close()
	fmt.Println(string(data), "wwwwwwwwwwwwwwwwwwwwwww")

}

func demo(w http.ResponseWriter, r *http.Request) {
	file := "/home/srlemon/Downloads/NVIDIA-Linux-x86_64-430.40.run"
	_, err := os.Stat(file)
	if ok := os.IsExist(err); !ok {
		http.NotFound(w, r)
	}
	http.ServeFile(w, r, file)

}
