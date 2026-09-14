package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)


func LoadConfig(cfg_filename string) (Config, error) {
	cfg := Config{}
	data, err := os.ReadFile(cfg_filename)
	if err != nil  {
		fmt.Println("Invalid or absent config file.")
		os.Exit(0)
	}
	if err := json.Unmarshal(data, &cfg); err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	return cfg, err
}
