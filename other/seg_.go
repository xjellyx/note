package main

import (
	"fmt"
	orm "github.com/suboat/sorm"

	"github.com/go-ego/gse"
)

var (
	text = "中华人民共和国中央人民政府，Handle word  segmentation results "

	seg gse.Segmenter
)

func cut() {
	hmm := seg.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = seg.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)

	hmm = seg.CutAll(text)
	fmt.Println("cut all: ", hmm)
}

func segCut() {
	// Text Segmentation
	tb := []byte(text)
	fmt.Println(seg.String(tb, true))

	segments := seg.Segment(tb)

	// Handle word segmentation results
	// Support for normal mode and search mode two participle,
	// see the comments in the code ToString function.
	// The search mode is mainly used to provide search engines
	// with as many keywords as possible
	fmt.Println(gse.ToString(segments, true))
}

func main() {
	// Loading the default dictionary
	seg.LoadDict()
	// Load the dictionary
	// seg.LoadDict("/data/allen/gocode/src/github.com/srlemon/note/dictionary.txt" + "/src/github.com/go-ego/gse/data/dict/dictionary.txt")

	cut()

	var (
		q     = orm.M{}
		orArr []interface{}
	)

	for _, _s := range []string{"aa", "bb"} {
		orArr = append(orArr, map[string]interface{}{
			"name": _s,
		})
	}
	q[orm.TagQueryKeyOr] = orArr

}
