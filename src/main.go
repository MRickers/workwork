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
	var restart bool
	var end bool
	var info bool

	flag.StringVar(&filename, "file", workSheetPath, "Specify worksheet.dat to use")
	flag.BoolVar(&start, "check-in", false, "Check in")
	flag.BoolVar(&restart, "restart", false, "Check in, discard previous check in")
	flag.BoolVar(&end, "check-out", false, "Check out")
	flag.BoolVar(&info, "info", false, "Shows info about current workday")

	flag.Parse()

	path, err := os.Executable()

	if err != nil {
		fmt.Println(err)
		return
	}
	if start {
		err := workwork.CheckInWorkDay(path)
		if err != nil {
			fmt.Println(err)
		}
	} else if restart {
		err := workwork.RestartWorkDay(path)
		if err != nil {
			fmt.Println(err)
		}
	} else if end {
		err := workwork.CheckOutWorkDay(path)
		if err != nil {
			fmt.Println(err)
		}
	} else if info {
		err := workwork.ShowInfo(path)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		flag.Usage()
	}
}
