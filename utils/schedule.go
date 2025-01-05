package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Junx27/ticket-booking/helper"
)

func UpdateSeatsStatus(scheduleID uint, seatsData map[string]string) error {
	jsonData, err := json.Marshal(seatsData)
	if err != nil {
		return fmt.Errorf("failed to marshal seats data: %v", err)
	}
	url := helper.GetUpdateScheduleSeats(scheduleID)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send PUT request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update seats, status code: %d", resp.StatusCode)
	}

	return nil
}
