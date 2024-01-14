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

	type PluginFactory = func(m map[string]any) plugins.Impl

	pluginFactory := map[string]PluginFactory{
		"discord": func(m map[string]any) plugins.Impl {
			return discord.From(m)
		},
		"local": func(m map[string]any) plugins.Impl {
			return local.From(m)
		},
	}

	var plugs []plugins.Impl
	var plugById = make(map[string]plugins.Impl)

	for _, p := range cfg.Configs {
		pluginName := p["plugin"].(string)
		id := p["id"].(string)

		makePlug, found := pluginFactory[pluginName]

		if !found {
			log.Fatal("Unsupported plugin:", pluginName)
		}

		plug := makePlug(p)

		plugById[id] = plug
		plugs = append(plugs, plug)
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
