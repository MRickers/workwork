package main

import (
	"flag"
	"fmt"
)

func main() {
	var filename string

	flag.StringVar(&filename, "file", "worksheet.dat", "Specify worksheet.dat to use")
	fmt.Println(filename)
}
