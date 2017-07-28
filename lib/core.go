package yarn

import "os/exec"

func getYarnCmd() *exec.Cmd {

	bin, err := exec.LookPath("yarn")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(bin)
	return cmd

}
