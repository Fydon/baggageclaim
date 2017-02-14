// +build !windows

package driver

import (
	"bytes"
	"fmt"
	"os/exec"
)

func (driver *NaiveDriver) GetVolumeSizeInBytes(path string) (int64, error) {
	stdout := &bytes.Buffer{}
	cmd := exec.Command("du", path)
	cmd.Stdout = stdout

	err := cmd.Run()
	if err != nil {
		return 0, err
	}

	var size int64
	_, err = fmt.Sscanf(stdout.String(), "%d", &size)
	if err != nil {
		return 0, err
	}

	return size, nil
}
