package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/cardigann/releaseinfo"
)

func main() {
	result, err := releaseinfo.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)
}
