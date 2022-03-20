package controller

import (
	"fmt"
	"strings"
	"time"

	"workwork/src/models"
)

type Loader interface {
	Load(filename string) (string, error)
	Save(filename string, data []byte) error
}

type Converter interface {
	Deserialize(data string) ([]models.WorkDay, error)
	Serialize(data []models.WorkDay) (string, error)
}

type PlainConverter struct {
}

func (converter *PlainConverter) Deserialize(data string) ([]models.WorkDay, error) {
	rows := strings.Split(data, "\r\n")

	workdays := []models.WorkDay{}
	for _, workday := range rows {
		currentDay := models.WorkDay{}
		data := strings.Split(workday, "\t")
		if len(data) != 2 {
			continue
		}
		date, err := time.Parse("01-02-2006", data[0])
		if err != nil {
			return workdays, err
		}
		currentDay.Date = date

		startEndTime := strings.Split(data[1], "-")

		begin, err := models.CreateDay(startEndTime[0])
		if err != nil {
			return workdays, err
		}
		currentDay.Begin = begin
		if len(startEndTime) == 2 {
			end, err := models.CreateDay(startEndTime[1])
			if err != nil {
				return workdays, err
			}
			currentDay.End = end
		}
		workdays = append(workdays, currentDay)
	}
	return workdays, nil
}

func (converter *PlainConverter) Serialize(data []models.WorkDay) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("workday empty")
	}
	workdays := "Date\tBegin-End\r\n"
	for _, workday := range data {
		workday_serialize := fmt.Sprintf("%s\t%02d:%02d-%02d:%02d\r\n",
			workday.Date.Format("01-02-2006"),
			workday.Begin.Hour,
			workday.Begin.Min,
			workday.End.Hour,
			workday.End.Min)

		workdays += workday_serialize
	}
	return workdays, nil
}
