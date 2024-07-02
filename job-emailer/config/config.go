package config

import (
	"encoding/csv"
	"log"
	"os"
)

type APIConfig struct {
	title   string
	baseURL string
	apiKey  string
}

func LoadConfig() []APIConfig {

	f, err := os.Open("auth.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var apiConfig []APIConfig

	for _, record := range data {
		var apiConfigLine APIConfig
		apiConfigLine.title = record[0]
		apiConfigLine.apiKey = record[1]
		apiConfigLine.baseURL = record[2]

		apiConfig = append(apiConfig, apiConfigLine)

	}

	return apiConfig

}
