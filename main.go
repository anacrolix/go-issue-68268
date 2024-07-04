package main

import (
	"fmt"
	"github.com/anacrolix/torrent/bencode"
	"github.com/google/go-cmp/cmp"
	"regexp"
)

func main() {
	if false {
		regexp.MustCompile(`[_\p{L}][_\p{L}\p{N}]*$`)
		fmt.Println("regexp compiled")
	}
	if false {
		cmp.Diff(1, 2)
		fmt.Println("diffed")
	}
	bencode.Marshal(nil)
	fmt.Println("main returning")
}
