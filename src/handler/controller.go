package controller

import (
	"fmt"
	"strings"
	"time"

	"workwork/src/models"

	"github.com/rs/zerolog/log"
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
			log.Warn().Err(err).Msg(fmt.Sprintf("could not parse date %s", data[0]))
			continue
		}
		currentDay.Date = date

		startEndTime := strings.Split(data[1], "-")

		begin, err := models.CreateDay(startEndTime[0])
		if err != nil {
			log.Warn().Err(err).Msg(fmt.Sprintf("could not parse day %s", startEndTime[0]))
			continue
		}
		currentDay.Begin = begin
		if len(startEndTime) == 2 {
			end, err := models.CreateDay(startEndTime[1])
			if err != nil {
				log.Warn().Err(err).Msg(fmt.Sprintf("could not parse day %s", startEndTime[1]))
				continue
			}
			currentDay.End = end
		}
		workdays = append(workdays, currentDay)
	}
	return workdays, nil
}

func (converter *PlainConverter) Serialize(data []models.WorkDay) (string, error) {
	return "", nil
}
