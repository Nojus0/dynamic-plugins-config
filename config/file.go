package config

type File struct {
	Configs []map[string]any `yaml:"configs"`
	Parts   []map[string]any `yaml:"parts"`
}
