package main

import (
	"fmt"
	"log"

	"github.com/Nojus0/configer/config"
	"github.com/Nojus0/configer/plugins"
	"github.com/Nojus0/configer/plugins/discord"
	"github.com/Nojus0/configer/plugins/local"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	var plugs []plugins.Impl
	var plugById = make(map[string]plugins.Impl)

	for _, p := range cfg.Configs {
		pluginName := p["plugin"].(string)
		id := p["id"].(string)

		switch pluginName {
		case "discord":
			plug := discord.From(p)
			plugById[id] = plug
			plugs = append(plugs, plug)
		case "local":
			plug := local.From(p)
			plugById[id] = plug
			plugs = append(plugs, local.From(p))
		default:
			log.Fatalln("Unsupported plugin:", pluginName)
		}

	}

	for _, plug := range plugs {
		plugHandler, id := plug.Info()
		fmt.Println("Plug INFO:\t", plugHandler, id)
	}

	for _, p := range cfg.Parts {
		id := p["id"].(string)
		partPlug := plugById[id]

		partPlug.HandlePart(p)

	}
}
