package yarn

import (
	"github.com/admmasters/easyutils"
)

// Install runs yarn install on the passed path
func Install(path string) {
	yarnCmd := getYarnCmd()
	easyutils.ExecuteCmd(yarnCmd, path)
}
