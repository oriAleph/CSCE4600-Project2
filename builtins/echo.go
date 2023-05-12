package builtins

import (
	"fmt"
)

func RepeatStuffs(args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("invalid argument count", ErrInvalidArgCount)
	default:
		fmt.Println(args[0:])
	}

	return nil
}
