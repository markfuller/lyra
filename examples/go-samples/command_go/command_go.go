package main

import (
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/lyraproj/servicesdk/lang/go/lyra"
)

type wordpressDeployment struct {
	Output string
}

func installWordpress() wordpressDeployment {
	log := hclog.Default()
	log.Debug("installWordpress entered")
	cmd := exec.Command("helm", "install", "--name", "wordpress", "stable/wordpress")
	out, err := cmd.Output()
	if err != nil {
		log.Debug("error running wordpress", "err", err)
		return wordpressDeployment{}
	}
	output := fmt.Sprintf("%s", out)

	return wordpressDeployment{Output: output}
}

func main() {

	// Lyra workflow
	lyra.Serve(`command_go`, nil, &lyra.Workflow{
		Return: struct {
			Output string
		}{},
		Steps: map[string]lyra.Step{
			`installWordpress`: &lyra.Action{
				Do: installWordpress,
			},
		},
	})
}
