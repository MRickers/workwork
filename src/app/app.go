package workwork

import (
	"fmt"
	"workwork/src/controller"
	"workwork/src/models"

	"github.com/rs/zerolog/log"
)

const workDayloadPath = "work.tmp"
const workSheetLoadPath = "worksheet.work"

func CheckInWorkDay() error {
	loader := controller.PlainLoader{}

	if loader.Exist(workDayloadPath) {
		return fmt.Errorf("already checked in")
	}

	workday := models.NewDay()
	converter := controller.PlainConverter{}
	var serialized = "Date\t\tBegin-End\r\n"

	serialized_day, err := converter.Serialize([]models.WorkDay{workday})
	if err != nil {
		log.Error().Err(err)
		return err
	}
	serialized += serialized_day
	err = loader.Save(workDayloadPath, []byte(serialized))

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
	err = loader.Delete(workDayloadPath)

	if err != nil {
		log.Error().Err(err)
		return err
	}
	// append to worksheet

	if loader.Exist(workSheetLoadPath) {
		err = loader.Append(workSheetLoadPath, []byte(serialized_day))

	} else {
		var serialized = "Date\t\tBegin-End\r\n" + serialized_day
		err = loader.Save(workSheetLoadPath, []byte(serialized))
	}

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}
