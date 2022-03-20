package main

import (
	"flag"
	"fmt"
	"os"

	workwork "workwork/src/app"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const workSheetPath = "worksheet.dat"

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	var filename string
	var start bool

	flag.StringVar(&filename, "file", workSheetPath, "Specify worksheet.dat to use")
	flag.BoolVar(&start, "start", false, "Command to get work begin timestamp")

	flag.Parse()

	if start {
		err := workwork.StartWorkDay()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		flag.Usage()
	}
}
