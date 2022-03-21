package workwork

import (
	"fmt"
	"workwork/src/controller"
	"workwork/src/models"

	"github.com/rs/zerolog/log"
)

const workDayloadPath = "work.tmp"

func CheckInWorkDay() error {
	loader := controller.PlainLoader{}

	if loader.Exist(workDayloadPath) {
		return fmt.Errorf("already checked in")
	}

	workday := models.NewDay()
	converter := controller.PlainConverter{}

	serialized_day, err := converter.Serialize([]models.WorkDay{workday})
	if err != nil {
		log.Error().Err(err)
		return err
	}
	err = loader.Save(workDayloadPath, []byte(serialized_day))

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}

func RestartWorkDay() error {
	loader := controller.PlainLoader{}

	workday := models.NewDay()
	converter := controller.PlainConverter{}

	serialized_day, err := converter.Serialize([]models.WorkDay{workday})
	if err != nil {
		log.Error().Err(err)
		return err
	}
	err = loader.Save(workDayloadPath, []byte(serialized_day))

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}

func CheckOutWorkDay() error {
	loader := controller.PlainLoader{}

	if !loader.Exist(workDayloadPath) {
		return fmt.Errorf("not checked in yet")
	}

	workday_byte, err := loader.Load(workDayloadPath)

	if err != nil {
		log.Error().Err(err)
		return err
	}
	converter := controller.PlainConverter{}

	workday, err := converter.Deserialize(string(workday_byte))

	if err != nil {
		log.Error().Err(err)
		return err
	}

	workday[0].Quit()

	serialized_day, err := converter.Serialize(workday)

	if err != nil {
		log.Error().Err(err)
		return err
	}
	err = loader.Save(workDayloadPath, []byte(serialized_day))

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}
