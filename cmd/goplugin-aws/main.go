package main

import (
	"flag"
	"github.com/hashicorp/go-hclog"
	"github.com/lyraproj/lyra/cmd/goplugin-aws/aws"
)

func main() {
	setLogLevel()
	aws.Start()
}
func setLogLevel() {
	debug := flag.Bool("debug", false, "debug logging")
	flag.Parse()

	if *debug {
		hclog.Default().SetLevel(hclog.Debug)
	}
}
