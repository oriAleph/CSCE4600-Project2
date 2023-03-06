package builtins

import (
	"errors"
	"fmt"
	"strings"
)

type AliasPair struct {
	alias    string
	original []string
}

var (
	ErrInvalidArg    = errors.New("invalid argument")
	aliasFormat      = "\nFormat: alias [option] [name]='[value]'\nExample: alias la='ls -a'\n"
	aliasOptions     = "\nFor list of options: alias help\n"
	aliasOptionsList = []string{"Print alias: alias -p [name]", "Remove alias: alias -r [name]",
		"Print all aliases: alias -p", "Remove all aliases: alias -ra", aliasFormat}
)

func Alias(aliasList []AliasPair, args ...string) error {

	argsSize := len(args)
	aliasSize := len(aliasList)

	switch argsSize {

	// Zero arguments
	case 0:
		fmt.Print(aliasFormat, aliasOptions)
		return nil

	// Check for one argument options
	case 1:
		if args[0] == "help" { // Options list and Format
			fmt.Print("\nOption List:\n")
			for _, h := range aliasOptionsList {
				fmt.Printf("%s\n", h)
			}
			fmt.Print(aliasFormat)
			return nil

		} else if args[0] == "-p" { // Print all aliases
			if aliasSize == 0 {
				fmt.Print("\nAlias list empty.\n")
			} else {
				fmt.Print("\nAlias List:\n")
				for _, p := range aliasList {
					fmt.Printf("%s = %s\n", p.alias, p.original)
				}
			}
			return nil

		} else if args[0] == "-ra" { // Remove all aliases
			if aliasSize == 0 {
				fmt.Print("\nAlias list empty.\n")
			} else {
				aliasList = nil
				fmt.Print("\nAll aliases removed.\n")
			}
			return nil

		}
	}

	// Multiple argument options
	if args[0] == "-p" { // Print alias
		for _, p := range aliasList {
			if p.alias == args[1] {
				fmt.Printf("\n%s=%s\n", p.alias, p.original)
				return nil
			}
		}
		return fmt.Errorf("%w: alias not found", ErrInvalidArg)

	} else if args[0] == "-r" { // Remove alias
		fmt.Print("\nDeleting alias...\n")
		for i, r := range aliasList {
			if r.alias == args[1] {
				aliasList = append(aliasList[:i], aliasList[i+1:]...)
				fmt.Printf("Alias deleted.\n")
				return nil
			}
		}
		return fmt.Errorf("%w: alias not found", ErrInvalidArg)

	}

	// Clean up arguments and append new alias
	// Format: alias [option] [name]='[value]'
	// Example: alias la='ls -a'
	// name, args := args[0], args[1:]
	if strings.Contains(args[0], "='") && strings.HasSuffix(args[argsSize-1], "'") {

		// Remove =' from first argument by splitting it
		var split = strings.Split(args[0], "='")

		if len(split) == 2 { // Just two from before and after ='

			args[0] = split[0]

			// Remove last single quote
			args[argsSize-1] = strings.TrimSuffix(args[argsSize-1], "'")

			// Combine arguments for full original command
			args = append(split[1:], args...)

			// Add alias pair to list
			aliasList = append(aliasList, AliasPair{alias: args[0], original: args[1:]})
			aliasSize = len(aliasList) //  new alias list size
			fmt.Printf("\nAdded alias: %s, original: %s\n", aliasList[aliasSize-1].alias, aliasList[aliasSize-1].original)
			return nil

		}

	} else {
		fmt.Print(aliasFormat, aliasOptions)
		return fmt.Errorf("%w: improper format", ErrInvalidArg)

	}

	fmt.Print(aliasFormat, aliasOptions)
	return fmt.Errorf("%w: logical error", ErrInvalidArg)
}
