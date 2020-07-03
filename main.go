package main

import (
	"flag"

	"log"

	"github.com/jaegertracing/jaeger/plugin/storage/grpc"
	"github.com/rubenvp8510/redbull/pkg/redbull"
)

//var logger = hclog.New(&hclog.LoggerOptions{
//Level:      hclog.Warn,
//Name:       "redbull",
//JSONFormat: true,
//})

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "A path to the plugin's configuration file")
	flag.Parse()

	plugin, err := redbull.NewRedBull(redbull.Config{})
	if err != nil {
		log.Panicln("error creating redbull", "err", err)
	}

	defer plugin.Close()

	grpc.Serve(plugin)
}
