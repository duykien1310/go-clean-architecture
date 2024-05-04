package util

import (
	"auth_service/config"
	"time"
)

func ConvertDateStringToPointer(dateData string) *time.Time {
	datePointer := &time.Time{}
	if dateData == "" {
		datePointer = nil
	}
	date, err := time.Parse(config.LAYOUT, dateData)
	if err != nil {
		datePointer = nil
	} else {
		datePointer = &date
	}
	return datePointer
}

func ConvertDateStringToTimeDate(dateData string) time.Time {
	date, _ := time.Parse(config.LAYOUT, dateData)
	return date
}

func ConvertLayoutDateTimeToDateWithoutTime(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
}
