package models

import (
	"fmt"
)

type Job struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Company     string `json:"company"`
	Location    string `json:"location"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func (j *Job) Validate() error {
	if j.ID == "" {
		return fmt.Errorf("ID is required")
	}
	if j.Title == "" {
		return fmt.Errorf("Title is required")
	}
	if j.Company == "" {
		return fmt.Errorf("Company is required")
	}
	if j.Location == "" {
		return fmt.Errorf("Location is required")
	}
	if j.Description == "" {
		return fmt.Errorf("Description is required")
	}
	if j.URL == "" {
		return fmt.Errorf("URL is required")
	}
	return nil
}
