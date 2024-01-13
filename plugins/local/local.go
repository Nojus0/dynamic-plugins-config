package local

import "fmt"

type Plugin struct {
	Id     string `yaml:"id"`
	Plugin string `yaml:"plugin"`
	Dir    string `yaml:"dir"`
}

type Part struct {
	Id   string `yaml:"id"`
	Path string `yaml:"path"`
}

func PartFrom(m map[string]any) *Part {
	id := m["id"].(string)
	path := m["path"].(string)

	return &Part{
		Id:   id,
		Path: path,
	}
}

func From(m map[string]any) *Plugin {

	id := m["id"].(string)
	plugin := m["plugin"].(string)
	dir := m["dir"].(string)

	return &Plugin{
		Id:     id,
		Plugin: plugin,
		Dir:    dir,
	}
}

func (p Plugin) Info() (string, string) { return p.Plugin, p.Id }
func (p Plugin) HandlePart(m map[string]any) error {
	part := PartFrom(m)

	fmt.Println("[LOCAL PLUGIN] IS EXECUTING PART", part.Id, "WITH DIR", part.Path)
	return nil
}
