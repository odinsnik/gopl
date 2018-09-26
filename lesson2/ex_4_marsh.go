package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Point struct {
	X,Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	filename := os.Args[1]
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(w)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(file, "%s\n", data)
	file.Close()

	fileIn, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var w2 Wheel
	b, err := ioutil.ReadAll(fileIn)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(b, &w2); err != nil {
		panic(err)
	}
	log.Printf("%#v\n", w2)
}

