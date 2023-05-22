package configuration

import (
	"RoRoDes/model"
	"encoding/json"
	"os"
)

func GetConfig(path string) model.Config {
	configInBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var config model.Config

	err = json.Unmarshal(configInBytes, &config)
	if err != nil {
		panic(err)
	}

	return config
}
