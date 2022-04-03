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

type WorkMonth struct {
	Date          time.Time
	TotalHours    uint32
	TotalMins     uint32
	OverTimeHours uint32
	OverTimeMins  uint32
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

// 20:00 04:30
func (work *WorkDay) Minus(other WorkDay) error {
	if work.Date.Before(other.Date) {
		return fmt.Errorf("other day greater this")
	}

	if work.Date.Year() > other.Date.Year() ||
		work.Date.Month() > other.Date.Month() ||
		work.Date.Day() > other.Date.Day() {

		fmt.Println("New day, year, month")
	} else { // same day
		work.Begin.Hour -= other.Begin.Hour
		work.Begin.Min -= other.Begin.Min
		if work.Begin.Min < 0 {
			work.Begin.Min += 60
			work.Begin.Hour -= 1
		}
	}
	return nil
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

func WorkSum(workSheet []WorkDay) {
	months := getMonths(workSheet)
	fmt.Print("Work statistics\r\n")
	for _, month := range months {
		fmt.Printf("%s %d\r\n%d:%02d hours Overtime: %d:%02d\r\n\r\n",
			month.Date.Month(),
			month.Date.Year(),
			month.TotalHours,
			month.TotalMins,
			month.OverTimeHours,
			month.OverTimeMins)
	}
	fmt.Printf("")
}

func getMonths(workSheet []WorkDay) []WorkMonth {
	const pause int = 30
	months := []WorkMonth{}

	if len(workSheet) == 0 {
		fmt.Print("empty worksheet")
		return nil
	}
	var lastMonth time.Month = workSheet[0].Date.Month()
	currentMonth := WorkMonth{Date: workSheet[0].Date, TotalHours: 0, TotalMins: 0}

	for _, day := range workSheet {
		if lastMonth != day.Date.Month() {
			lastMonth = day.Date.Month()
			// minuten in stunden umrechnen
			mins := currentMonth.TotalMins % 60
			hours := currentMonth.TotalMins / 60
			currentMonth.TotalHours += hours
			currentMonth.TotalMins = mins

			mins = currentMonth.OverTimeMins % 60
			hours = currentMonth.OverTimeMins / 60
			currentMonth.OverTimeHours += hours
			currentMonth.OverTimeMins = mins

			months = append(months, currentMonth)
			currentMonth = WorkMonth{Date: day.Date, TotalHours: 0, TotalMins: 0}
		}
		workedHours := day.End.Hour - day.Begin.Hour
		workedMins := day.End.Min - day.Begin.Min - pause
		fmt.Printf("\n\nHours: %d Mins: %d", workedHours, workedMins)
		if workedHours < 0 {
			fmt.Printf("invalid workedHours %d", workedHours)
			continue
		}

		if workedMins < 0 && workedHours >= 1 {
			workedHours -= 1
			workedMins += 60
		}

		// 0810 1603
		if workedHours > 8 || (workedHours == 8 && workedMins > 0) {
			currentMonth.OverTimeHours += uint32(workedHours - 8)
			currentMonth.OverTimeMins += uint32(workedMins)
		}

		if workedHours >= 0 && workedMins >= 0 {
			currentMonth.TotalHours += uint32(workedHours)
			currentMonth.TotalMins += uint32(workedMins)
		}
	}

	mins := currentMonth.TotalMins % 60
	hours := currentMonth.TotalMins / 60
	currentMonth.TotalHours += hours
	currentMonth.TotalMins = mins

	mins = currentMonth.OverTimeMins % 60
	hours = currentMonth.OverTimeMins / 60
	currentMonth.OverTimeHours += hours
	currentMonth.OverTimeMins = mins

	months = append(months, currentMonth)

	return months
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
