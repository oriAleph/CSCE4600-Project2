package builtins

import(
	"fmt"
	"os"
)

func mkdir(args ...string) error {
	switch len(args) 
	{
	case 1:
		err := os.mkdir(args[0], 0755)
		if err != nil {
			fmt.PrintLn(err)
		}
		else {
			println("directory successfully created")
			return err
		}
	default:
		return fmt.Errorf("%w, at least one argument expected", ErrInvalidArgCount)
	}
}