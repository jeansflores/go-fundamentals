package main

import (
	data "example.com/platform/data"
	sampledata "example.com/platform/sample_data"
)

var courses []data.Printable

func main() {
	courses = sampledata.GetData()
	for _, c := range courses {
		println(c.Print())
	}
}
