package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
)

func RepeatStuffs(args ...string) error {
	switch len(args) {
	case 0: 
			return fmt.Errorf(ErrInvalidArgCount)
	default:
		return fmt.Println(args[0:])
	}
}
