//Code by Kennedy Boynton KLB0583
//Group 18 5/10/23 CSCE 4600.003

package builtins

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetDirectory(w io.Writer, args ...string) error {

	ErrInvalidArgCount = errors.New("invalid argument count")

	switch len(args) {
	case 0:
		dir, err := os.Getwd()
		//no arguments
		if err != nil {
			return err
		}

		//fmt.Println(dir)
		toShow := make([]string, 0)
		toShow = append(toShow, dir)

		_, err2 := fmt.Fprintln(w, strings.Join(toShow, "\n"))

		return err2

	case 1:
		if args[0] == "-L" || args[0] == "-l" {
			//logical link get
			dir, err := os.Getwd()

			if err != nil {
				return err
			}

			//fmt.Println(dir)
			toShow := make([]string, 0)
			toShow = append(toShow, dir)

			_, err2 := fmt.Fprintln(w, strings.Join(toShow, "\n"))

			return err2
		}

		if args[0] == "-P" || args[0] == "-p" {
			//physical link get
			target, err := os.Readlink("/proc/self/cwd")

			if err != nil {
				return err
			}

			//fmt.Println(dir)
			toShow := make([]string, 0)
			toShow = append(toShow, target)

			_, err2 := fmt.Fprintln(w, strings.Join(toShow, "\n"))

			return err2
		}

		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCount)

	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCount)
	}
}
