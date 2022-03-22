package controller

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"workwork/src/models"
)

type FileLoader interface {
	Load(filename string) ([]byte, error)
	Save(filename string, data []byte) error
	Append(filename string, data []byte) error
	Delete(filename string) error
	Exist(filename string) bool
}

type Converter interface {
	Deserialize(data string) ([]models.WorkDay, error)
	Serialize(data []models.WorkDay) (string, error)
}

type PlainLoader struct{}

func (loader *PlainLoader) Load(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func (loader *PlainLoader) Save(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)

	if err != nil {
		return err
	}
	return nil
}

func (loader *PlainLoader) Append(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)

	if err != nil {
		return err
	}
	return nil
}

func (loader *PlainLoader) Delete(filename string) error {
	return os.Remove(filename)
}

func (loader *PlainLoader) Exist(filename string) bool {
	_, err := os.Stat(filename)
	return !errors.Is(err, os.ErrNotExist)
}

type PlainConverter struct {
}

func (converter *PlainConverter) Deserialize(data string) ([]models.WorkDay, error) {
	rows := strings.Split(data, "\r\n")

	workdays := []models.WorkDay{}
	for _, workday := range rows {
		currentDay := models.WorkDay{}
		data := strings.Split(workday, "\t")
		if len(data) != 3 {
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
		currentDay.Pause, err = strconv.Atoi(data[2])
		if err != nil {
			return workdays, err
		}
		workdays = append(workdays, currentDay)
	}
	return workdays, nil
}

func (converter *PlainConverter) Serialize(data []models.WorkDay) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("workday empty")
	}
	workdays := ""
	for _, workday := range data {
		workday_serialize := fmt.Sprintf("%s\t%02d:%02d-%02d:%02d\t%d\r\n",
			workday.Date.Format("01-02-2006"),
			workday.Begin.Hour,
			workday.Begin.Min,
			workday.End.Hour,
			workday.End.Min,
			workday.Pause)

		workdays += workday_serialize
	}
	return workdays, nil
}
