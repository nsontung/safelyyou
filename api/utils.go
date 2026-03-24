package api

import (
	"encoding/csv"
	"os"
)

func LoadDeviceIDsFromCSV(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var deviceIDs []string
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records[1:] { // Skip header
		if len(record) > 0 {
			deviceIDs = append(deviceIDs, record[0])
		}
	}
	return deviceIDs, nil

}
