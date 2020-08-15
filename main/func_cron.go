package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/olongfen/note/log"
	"os"
)

var (
	defaultBufSize int64 = 4096
)

func main() {
	//f := "/data/gocode/src/github.com/olongfen/note/demo.sh"
	//lines, _ := tail(f, 10000)
	m := make([]byte, 0)
	m, _ = base64.StdEncoding.DecodeString("RXhlY3V0ZSA6L3Vzci9iaW4vcHl0aG9uMyAuL3B1YmxpYy9zY3JpcHQvZ2VuX3NpZ25hbC5weSBmYWlsZW" +
		"Qgd2l0aCBlcnJvcjpleGl0IHN0YXR1cyAzLCBvdXRwdXQ6IC92YXIvd3d3L3Jkcy9iaW4vd3RpLXNpZ25hbAplbnRlcmVkIHdvcmsgZGlyOiAvdG" +
		"1wLzIwMjAtMDgtMDcKcHJlcGFyaW5nIG9yaWdpbiBkYXRhIGFzIHN0b3JhZ2Vfc3Rhcndpel9vcmlnaW5hbC50eHQgLi4uCkVycm9yOiAgKHBzeW" +
		"NvcGcyLk9wZXJhdGlvbmFsRXJyb3IpIGNvdWxkIG5vdCBjb25uZWN0IHRvIHNlcnZlcjogQ29ubmVjdGlvbiByZWZ1c2VkCglJcyB0aGUgc2VydmVyI" +
		"HJ1bm5pbmcgb24gaG9zdCAiMTI3LjAuMC4xIiBhbmQgYWNjZXB0aW5nCglUQ1AvSVAgY29ubmVjdGlvbnMgb24gcG9ydCA1NDMyPwoKKEJhY2tncm9" +
		"1bmQgb24gdGhpcyBlcnJvciBhdDogaHR0cDovL3NxbGFsY2hlLm1lL2UvMTMvZTNxOCkKZ2F0aGVyaW5nIGRhdGEgLi4uCkVSUk9SOiBmYWlsZ" +
		"WQgdG8gZ2F0aGVyaW5nIGRhdGEuCg==")
	fmt.Println(string(m))
	//f_, _ := os.Open(f)
	//d, _ := f_.Stat()
	//s := d.Size()
	//f_.Seek(s-defaultBufSize, os.SEEK_SET)
	//b := make([]byte, defaultBufSize)
	//dd, _ := f_.Read(b)
	//fmt.Println("sdsdsad", string(b), dd)

}
