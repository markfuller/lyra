
package main

import (
	"strings"
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/lyraproj/servicesdk/lang/go/lyra"
)

type helmIn struct {
	Name string
	Chart string
	Values []string
	Namespace *string
}

type helmOut struct {
	Output string
}

func helmInstall(in helmIn) helmOut {
	log := hclog.Default()
	namespace := "default"
	if in.Namespace != nil {
		namespace = *in.Namespace
	}
	args := []string {
		"install", 
		"--namespace",
		namespace,
		"--name", 
		in.Name,
		in.Chart,

	}
	if len(in.Values) > 0 {
		args = append(args, "--set")

		//HACK: unsure why this Replace is needed but strings.Join seems to add space instead
		x := strings.Replace(strings.Join(in.Values, ",")," ", ",", -1)
		// x := strings.Join(in.Values, ",")
		args = append(args, x)
	}
	cmd := exec.Command("helm", args...)

	log.Debug("about to run command", "cmd", cmd)

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
