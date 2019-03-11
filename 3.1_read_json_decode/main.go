package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Database struct {
		SQL      string `json:"sql"`
		Hostip   string `json:"hostip"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Env     string `json:"env"`
	Timeout string `json:"timeout"`
}

func main() {
	file := "config.json"
	config, _ := LoadConfig(file)
	fmt.Println(config)
}

func LoadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	// Always defer file-open handlers
	//defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&config); err != nil {
		return config, err
	}
	return config, err
}
