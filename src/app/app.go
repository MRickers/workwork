package workwork

import (
	"fmt"
	"workwork/src/controller"
	"workwork/src/models"

	"github.com/rs/zerolog/log"
)

const loadPath = "work.tmp"

func StartWorkDay() error {
	loader := controller.PlainLoader{}

	if loader.Exist(loadPath) {
		return fmt.Errorf("day already started")
	}

	workday := models.NewDay()
	converter := controller.PlainConverter{}

	serialized_day, err := converter.Serialize([]models.WorkDay{workday})
	if err != nil {
		log.Error().Err(err)
		return err
	}
	err = loader.Save(loadPath, []byte(serialized_day))

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}
