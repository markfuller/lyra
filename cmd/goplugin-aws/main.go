package main

import (
	"flag"
	"github.com/hashicorp/go-hclog"
	"github.com/leonelquinteros/gotext"
	"github.com/lyraproj/lyra/cmd/goplugin-aws/aws"
)

func main() {
	setLogLevel()
	aws.Start()
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
			hclog.Default().Error("set log level to ", "logLevel", logLevel)
		}
	}

}
