package workwork

import (
	"fmt"
	"workwork/src/controller"
	"workwork/src/helper"
	"workwork/src/models"

	"github.com/rs/zerolog/log"
)

const workDayloadPath = "work.tmp"
const workSheetLoadPath = "worksheet.work"

func CheckInWorkDay(exePath string) error {
	var absDayLoadPath = helper.StripExeName(exePath) + workDayloadPath

	loader := controller.PlainLoader{}

	if loader.Exist(absDayLoadPath) {
		return fmt.Errorf("already checked in")
	}

	workday := models.NewDay()
	converter := controller.PlainConverter{}
	var serialized = "Date\t\tBegin-End\tPause\r\n"

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
	var absDayLoadPath = helper.StripExeName(exePath) + workDayloadPath

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
	var absDayLoadPath = helper.StripExeName(exePath) + workDayloadPath
	var absSheetLoadPath = helper.StripExeName(exePath) + workSheetLoadPath

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
	workday[0].Pause = 30
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
		var serialized = "Date\t\tBegin-End\tPause\r\n" + serialized_day
		err = loader.Save(absSheetLoadPath, []byte(serialized))
	}

	if err != nil {
		log.Error().Err(err)
		return err
	}
	return nil
}

func ShowInfo(exePath string) error {
	var absDayLoadPath = helper.StripExeName(exePath) + workDayloadPath

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
	day := workday[0]

	if err != nil {
		log.Error().Err(err)
		return err
	}
	var timeWorked = models.NewDay()

	err = timeWorked.Minus(day)
	if err != nil {
		return err
	}

	fmt.Printf("%s\r\nChecked in: %02d:%02d\r\nTime worked: %02d:%02d",
		day.Date.Format("01-02-2006"),
		day.Begin.Hour,
		day.Begin.Min,
		timeWorked.Begin.Hour,
		timeWorked.Begin.Min)

	return nil
}

func ShowOverallInfo(exePath string) error {
	var absSheetLoadPath = helper.StripExeName(exePath) + workSheetLoadPath

	loader := controller.PlainLoader{}

	if !loader.Exist(absSheetLoadPath) {
		return fmt.Errorf("worksheet found")
	}

	workSheet_byte, err := loader.Load(absSheetLoadPath)

	if err != nil {
		return fmt.Errorf("loading worksheet failed: %s", err.Error())
	}

	converter := controller.PlainConverter{}

	workSheet, err := converter.Deserialize(string(workSheet_byte))

	if err != nil {
		return fmt.Errorf("deserialzing worksheet failed: %s", err.Error())
	}

	models.WorkSum(workSheet)

	return nil
}
