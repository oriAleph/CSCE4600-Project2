package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/oriAleph/CSCE4600-Project2/builtins"
)

type (
	Builtin struct {
		command  string
		function error
	}
	AliasPair struct {
		alias    string
		original []string
	}
)

var aliasList []AliasPair

func main() {
	exit := make(chan struct{}, 2) // buffer this so there's no deadlock.
	runLoop(os.Stdin, os.Stdout, os.Stderr, exit)
}

func runLoop(r io.Reader, w, errW io.Writer, exit chan struct{}) {
	var (
		input    string
		err      error
		readLoop = bufio.NewReader(r)
	)
	for {
		select {
		case <-exit:
			_, _ = fmt.Fprintln(w, "exiting gracefully...")
			return
		default:
			if err := printPrompt(w); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if input, err = readLoop.ReadString('\n'); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if err = handleInput(w, input, exit); err != nil {
				_, _ = fmt.Fprintln(errW, err)
			}
		}
	}
}

func printPrompt(w io.Writer) error {
	// Get current user.
	// Don't prematurely memoize this because it might change due to `su`?
	u, err := user.Current()
	if err != nil {
		return err
	}
	// Get current working directory.
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	// /home/User [Username] $
	_, err = fmt.Fprintf(w, "%v [%v] $ ", wd, u.Username)

	return err
}

func handleInput(w io.Writer, input string, exit chan<- struct{}) error {
	// Remove trailing spaces.
	input = strings.TrimSpace(input)

	// Split the input separate the command name and the command arguments.
	args := strings.Split(input, " ")
	name, args := args[0], args[1:]

	// Fill in builtin command list
	// New builtin commands should be added here
	var builtinList = []Builtin{
		{command: "cd", function: builtins.ChangeDirectory(args...)},
		{command: "env", function: builtins.EnvironmentVariables(w, args...)},
		{command: "alias", function: builtins.Alias(aliasList, args...)},
	}

	// Check for built-in commands
	if name == "exit" {
		exit <- struct{}{}
		return nil
	}
	for _, b := range builtinList {
		if name == b.command {
			return b.function
		}
	}

	// Check for alias commamds
	for _, a := range aliasList {
		if name == a.alias {
			return executeCommand(a.original[0], a.original[1:]...)
		}
	}

	return executeCommand(name, args...)
}

func executeCommand(name string, arg ...string) error {
	// Otherwise prep the command
	cmd := exec.Command(name, arg...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
