package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const (
	AudioSamplingRateMP3  = "16000"
	AudioBitRate          = "12.2k" // in Hz
	NumberOfAudioChannels = "1"
	AudioSamplingRateAMR  = "8000"
)

func main() {
	err := converToMP3("/data/allen/gocode/src/github.com/srlemon/note/other_/aaa.mp3")
	fmt.Println(err)
	now := time.Now()
	t := now.Add(time.Minute * -1)
	fmt.Println(t.String(), now.String())
}

func converToMP3(filename string) (err error) {
	var (
		toName   string
		fromName = filename
	)

	if len(filename) == 0 {
		err = errors.New("参数未填")
		return
	}
	toName = filename
	a := strings.Split(toName, ".mp3")
	fmt.Println(a)
	comm := exec.Command("ffmpeg", "-i", fromName, "-ar", AudioSamplingRateMP3, toName+".wav")
	fmt.Println(comm)
	if err = comm.Run(); err != nil {
		return
	}
	return
}
