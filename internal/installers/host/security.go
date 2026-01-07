package host

import (
	"os"
	"os/exec"
)

func securePath(path string) error {
	if err := os.Chown(path, 0, 0); err != nil {
		return err
	}
	if err := os.Chmod(path, 0750); err != nil {
		return err
	}
	return exec.Command("chattr", "-R", "+i", path).Run()
}

