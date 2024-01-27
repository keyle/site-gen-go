package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func LoadSettings() (*Settings, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	settingsFile := ".config/site-gen/settings.json"
	fullPath := filepath.Join(homeDir, settingsFile)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	var settings Settings
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}
