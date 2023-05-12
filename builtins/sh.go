//Writen by Saul Garay

package builtins

import (
	"fmt"
	"os"
	"os/exec"
)

func SH(args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("%w: expected more args", ErrInvalidArgCount)
	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}
