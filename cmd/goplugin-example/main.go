package main

import (
	"flag"
	"github.com/hashicorp/go-hclog"
	"github.com/leonelquinteros/gotext"
	"github.com/lyraproj/lyra/cmd/goplugin-example/example"
	"log"
)

func main() {
	setLogLevel()
	example.Start()
}

func setLogLevel() {
	logLevel := flag.String("loglevel", "", gotext.Get("Set log level to error, warn, info, debug or trace"))
	flag.Parse()

	hclog.Default().SetLevel(3)
	hclog.Default().Error("set log level to INFO", "logLevel", logLevel)

	if logLevel != nil && len(*logLevel) > 0 {
		ll := hclog.LevelFromString(*logLevel)
		if ll != hclog.NoLevel {
			hclog.Default().SetLevel(ll)
			hclog.DefaultOptions.Level = ll
		}
	}
	hclog.Default().Warn("warning")
	hclog.Default().Info("info")
	hclog.Default().Debug("debug")
	hclog.Default().Error("fatal")
	hclog.Default().Trace("trace")
	log.Println("THIS IS A LOG MESSAGE")
}
