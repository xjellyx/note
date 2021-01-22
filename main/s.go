package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	f, _ := ioutil.ReadFile("./d.json")
	var (
		data = make(map[string]interface{})
	)
	json.Unmarshal(f, &data)
	fmt.Println(len(data))
	ff, _ := os.Open("./dd.txt")
	defer ff.Close()
	br := bufio.NewReader(ff)
	count := 0
	for {
		s, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if strings.Contains(string(s), "Tile") {
			count++
		}
	}
	fmt.Println(count)

}

func ee() {
	var (
		arrStr = []string{"a\\b\\c", "a\\d\\e", "b\\cst", "d\\"}
	)
	for _, v := range arrStr {
		l := strings.Split(v, "\\")
		var (
			space = ""
		)
		for _i, _v := range l {
			if _i >= 1 {
				space += " "
			}
			fmt.Printf("%s%s\n", space, _v)
		}
	}

}

/*
201710/t4
201708/t6
201409/t1.t2
201411/t3
201401/t6
201907/t1
201911/t3.t5
201908/t4
201909/t4
201903/t6
201801/t4,t6
201210/t4
201504/t4
201303/t3
201301/t4
202001/t1
202004/t1
*/
