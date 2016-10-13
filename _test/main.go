package main

import (
	"log"

	"github.com/dlclark/regexp2"
)

func main() {
	target := "The.100000.Dollar.Pyramid.2016.S01E05.720p.HDTV.x264-W4F"

	re := regexp2.MustCompile(
		`^(?<title>.+?)(?:\W+(?:(?:Part\W?|(?<!\d+\W+)e)(?<episode>\d{1,2}(?!\d+)))+)`,
		regexp2.IgnoreCase|regexp2.Compiled)

	match, err := re.FindStringMatch(target)
	if err != nil {
		log.Fatal(err)
	}

	if match == nil {
		log.Fatal("No Match, correct")
	}

	log.Printf("Matched")
	for _, group := range match.Groups() {
		log.Fatalf("Group %s [%d]: %s", group.Name, group.Length, group.Capture.String())
	}

}
