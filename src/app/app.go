package workwork

import (
	"fmt"
	"workwork/src/controller"
	"workwork/src/models"

	"github.com/rs/zerolog/log"
)

const workDayloadPath = "work.tmp"
const workSheetLoadPath = "worksheet.work"

func CheckInWorkDay(exePath string) error {
	var absDayLoadPath = exePath + "\\" + workDayloadPath

	loader := controller.PlainLoader{}

	if loader.Exist(absDayLoadPath) {
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

	err = loader.Save(absDayLoadPath, []byte(serialized))

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}

func RestartWorkDay(exePath string) error {
	var absDayLoadPath = exePath + "\\" + workDayloadPath

	loader := controller.PlainLoader{}

	workday := models.NewDay()
	converter := controller.PlainConverter{}

	serialized_day, err := converter.Serialize([]models.WorkDay{workday})
	if err != nil {
		log.Error().Err(err)
		return err
	}
	err = loader.Save(absDayLoadPath, []byte(serialized_day))

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}

func CheckOutWorkDay(exePath string) error {
	var absDayLoadPath = exePath + "\\" + workDayloadPath
	var absSheetLoadPath = exePath + "\\" + workSheetLoadPath

	loader := controller.PlainLoader{}

	if !loader.Exist(absDayLoadPath) {
		return fmt.Errorf("not checked in yet")
	}

	workday_byte, err := loader.Load(absDayLoadPath)

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
	err = loader.Delete(absDayLoadPath)

	if err != nil {
		log.Error().Err(err)
		return err
	}
	// append to worksheet

	if loader.Exist(absSheetLoadPath) {
		err = loader.Append(absSheetLoadPath, []byte(serialized_day))

	} else {
		var serialized = "Date\t\tBegin-End\r\n" + serialized_day
		err = loader.Save(absSheetLoadPath, []byte(serialized))
	}

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}
