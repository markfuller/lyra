package main

import (
	"fmt"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/lyraproj/servicesdk/lang/go/lyra"
)

type cmdIn struct {
	Command string
}

type cmdOut struct {
	Output string
}

func runCommand(in cmdIn) cmdOut {
	log := hclog.Default()

	//run bash -c <in.Command> in a not very pretty way
	args := append([]string{"-c", in.Command})
	cmd := exec.Command("bash", args...)

	log.Debug("about to run command", "cmd", cmd)

	out, err := cmd.CombinedOutput()
	output := fmt.Sprintf("%s", out)
	if err != nil {
		panic(fmt.Errorf("error running cmd %v \n error is %v \n output is %v", cmd, err, output))
	}

	return cmdOut{Output: output}
}

func main() {
	lyra.Serve(`command_go`, nil, &lyra.Action{Do: runCommand})
}
