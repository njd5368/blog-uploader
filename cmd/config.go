package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const CONFIG_FILE_PATH = "/.config/blog-uploader/config.json"

type Config struct {
	APIURL string `json="apiurl"`
}

func ReadConfigFile() (*Config, error) {
	h, err := os.UserHomeDir()

	if err != nil {
		fmt.Print("Problem getting home directory: " + err.Error())
		return nil, err
	}

	b, err := ioutil.ReadFile(h + CONFIG_FILE_PATH)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Print("Config file does not exist at ~" + CONFIG_FILE_PATH + "\n" + err.Error())
		return nil, err
	}
	if err != nil {
		fmt.Print("Problem opening ~" + CONFIG_FILE_PATH + "\n" + err.Error())
		return nil, err
	}

	if err != nil {
		fmt.Print("Problem reading config file\n" + err.Error())
		return nil, err
	}

	var c Config
	err = json.Unmarshal(b, &c)

	if err != nil {
		fmt.Print("Problem unmarshaling data from config\n" + err.Error())
		return nil, err
	}

	return &c, nil
}