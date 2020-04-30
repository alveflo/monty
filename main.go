package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// DataPoint represents an Alpha order chart data point
type DataPoint struct {
	Timestamp int16   `json:"timestamp"`
	Price     float32 `json:"price"`
}

func readFile(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	var arr []DataPoint
	byteValue, _ := ioutil.ReadAll(jsonFile)

	_ = json.Unmarshal(byteValue, &arr)

	defer jsonFile.Close()
}

func main() {
	var filepath string

	flag.StringVar(&filepath, "filepath", "", "File to read")
	flag.Parse()

	if len(filepath) == 0 {
		fmt.Println("Please provide a filepath. (Hint: pass --filepath argument)")
	} else {
		readFile(filepath)
	}
}
