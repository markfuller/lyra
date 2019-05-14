
package main

import (
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/lyraproj/servicesdk/lang/go/lyra"
)

type helmIn struct {
	Name string
	Chart string	
	Namespace *string
}

type helmOut struct {
	Output string
}

func helmInstall(in helmIn) helmOut {
	log := hclog.Default()
	log.Debug("helmInstall entered", "in", in)
	namespace := "default"
	if in.Namespace != nil {
		namespace = *in.Namespace
	}
	cmd := exec.Command("helm", "install", 
		"--namespace",
		namespace,
		"--name", 
		in.Name,
		in.Chart)
	out, err := cmd.Output()
	output := fmt.Sprintf("%s", out)
	if err != nil {
		log.Debug("error running helm", "err", err, "cmd", cmd, "output", output)
		return helmOut{}
	}
	

	return helmOut{Output: output}
}

func main() {

	// Lyra workflow
	lyra.Serve(`helm_go`, nil, &lyra.Workflow{
		Parameters: helmIn{},
		Return: struct {
			Output string
		}{},
		Steps: map[string]lyra.Step{
			`helmInstall`: &lyra.Action{
				Do: helmInstall,
			},
		},
	})
}
