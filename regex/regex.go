package main

import (
	"fmt"
	"regexp"
)

func main() {
	const sample = `12,456.00 บาท`

	r, _ := regexp.Compile(`(^|,|บาท|$)|\s`)
	s := r.ReplaceAllString(sample, "")
	fmt.Println(len(s))
}
