package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

const configName = "config.yaml"

func (f File) Save() error {

	out, err := yaml.Marshal(f)
	if err != nil {
		return err
	}

	file, err := os.Create(configName)
	defer file.Close()
	if err != nil {
		return err
	}

	file.Write(out)
	return nil
}

func Load() (*File, error) {

	file, err := os.Open(configName)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	var out File

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
