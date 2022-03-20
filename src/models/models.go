package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Day struct {
	Hour int
	Min  int
}

type WorkDay struct {
	Date  time.Time
	Begin Day
	End   Day
	Pause int
}

func (work *WorkDay) Start() {
	work.Date = time.Now()
	work.Begin.Hour = work.Date.Hour()
	work.Begin.Min = work.Date.Minute()
}

func (work *WorkDay) Quit() {
	end := time.Now()
	work.End.Hour = end.Hour()
	work.End.Min = end.Minute()
}

func NewDay() WorkDay {
	day := WorkDay{}
	day.Start()
	return day
}

// Creates and validates day struct from day string hh:mm
func CreateDay(day string) (Day, error) {
	startEndTime := strings.Split(day, ":")

	if len(startEndTime) != 2 {
		return Day{}, fmt.Errorf("could not parse day string %s", day)
	}
	hour, err := strconv.Atoi(startEndTime[0])
	if err != nil {
		return Day{}, err
	}
	if !validHour(hour) {
		return Day{}, fmt.Errorf("invalid hour %d", hour)
	}

	if len(startEndTime) == 2 {
		minute, err := strconv.Atoi(startEndTime[1])
		if err != nil {
			return Day{}, err
		}
		if !validMinute(minute) {
			return Day{}, fmt.Errorf("invalid minute %d", minute)
		}
		return Day{Hour: hour, Min: minute}, nil
	}
	return Day{Hour: hour, Min: 0}, nil
}

func validHour(hour int) bool {
	if hour <= 24 && hour >= 0 {
		return true
	}
	return false
}

func validMinute(minute int) bool {
	if minute <= 59 && minute >= 0 {
		return true
	}
	return false
}
