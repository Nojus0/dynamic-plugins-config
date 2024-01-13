package discord

import "fmt"

type Part struct {
	Id         string `yaml:"id"`
	Attachment string `yaml:"attachment"`
}

func PartFrom(m map[string]any) *Part {
	id := m["id"].(string)
	attachment := m["attachment"].(string)

	return &Part{
		Id:         id,
		Attachment: attachment,
	}
}

type Plugin struct {
	Id        string `yaml:"id"`
	Plugin    string `yaml:"plugin"`
	AuthToken string `yaml:"auth_token"`
	ChannelId string `yaml:"channel_id"`
}

func From(m map[string]any) *Plugin {

	id := m["id"].(string)
	plugin := m["plugin"].(string)
	authToken := m["auth_token"].(string)
	channelId := m["channel_id"].(string)

	return &Plugin{
		Id:        id,
		Plugin:    plugin,
		AuthToken: authToken,
		ChannelId: channelId,
	}
}

func (p Plugin) HandlePart(m map[string]any) error {

	part := PartFrom(m)

	fmt.Println("PLUGIN DISCORD IS EXECUTING", part.Id, "WITH ATTACHMENT", part.Attachment)
	return nil
}

func (p Plugin) Info() (string, string) { return p.Plugin, p.Id }
