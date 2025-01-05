package helper

import (
	"fmt"
)

const BaseURL = "http://localhost:8080"

func GetUpdateScheduleSeats(scheduleID uint) string {
	return fmt.Sprintf("%s/schedules/%d/seats", BaseURL, scheduleID)
}
