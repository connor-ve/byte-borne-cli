package lib

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	PreviousTime string `yaml:"previous_time"`
}

// Function to read the previous time from the YAML file
func readPreviousTime() (string, error) {
	filePath := "C:\\Users\\crvan\\dark_terminal\\database\\time_db.yaml"
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}

	return config.PreviousTime, nil
}

// Function to calculate the time in seconds from the current time
func calculateTimeDifference(previousTime string) (int64, error) {
	prevTime, err := time.Parse(time.RFC3339, previousTime)
	if err != nil {
		return 0, err
	}

	currentTime := time.Now()
	diff := currentTime.Sub(prevTime)

	return int64(diff.Seconds()), nil
}

// Function to write the updated time back to the YAML file
func writeCurrentTime() error {
	filePath := "C:\\Users\\crvan\\dark_terminal\\database\\time_db.yaml"
	currentTime := time.Now().Format(time.RFC3339)

	config := Config{
		PreviousTime: currentTime,
	}

	newData, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, newData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetAfkTime() int64 {
	log.Println("Starting GetAfkTime...")

	// Read the previous time from the YAML file
	previousTime, err := readPreviousTime()
	if err != nil {
		log.Fatalf("Error reading previous time: %v", err)
	}

	timeDiff, err := calculateTimeDifference(previousTime)
	if err != nil {
		log.Fatalf("Error calculating time difference: %v", err)
	}

	// Write the current time back to the YAML file
	err = writeCurrentTime()
	if err != nil {
		log.Fatalf("Error writing current time: %v", err)
	}

	log.Println("GetAfkTime completed successfully.")
	return timeDiff
}
