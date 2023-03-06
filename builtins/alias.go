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
	aliasList []AliasPair
	alias     []string   //transfer values over since you can't pass structs
	original  [][]string //transfer values over since you can't pass structs
)

func Alias(args ...string) ([]string, [][]string, error) {

	argsSize := len(args)
	aliasSize := len(aliasList)

	switch argsSize {

	// Zero arguments
	case 0:
		fmt.Print(aliasFormat, aliasOptions)
		return alias, original, nil

	// Check for one argument options
	case 1:
		if args[0] == "help" { // Options list and Format
			fmt.Print("\nOption List:\n")
			for _, h := range aliasOptionsList {
				fmt.Printf("%s\n", h)
			}
			fmt.Print(aliasFormat)
			return alias, original, nil

		} else if args[0] == "-p" { // Print all aliases
			if aliasSize == 0 {
				fmt.Print("\nAlias list empty.\n")
			} else {
				fmt.Print("\nAlias List:\n")
				for _, p := range aliasList {
					fmt.Printf("%s = %s\n", p.alias, p.original)
				}
			}
			return alias, original, nil

		} else if args[0] == "-ra" { // Remove all aliases
			if aliasSize == 0 {
				fmt.Print("\nAlias list empty.\n")
			} else {
				aliasList = nil
				alias = nil
				original = nil
				fmt.Print("\nAll aliases removed.\n")
			}
			return alias, original, nil

		}
	}

	// Multiple argument options
	if args[0] == "-p" { // Print alias
		for _, p := range aliasList {
			if p.alias == args[1] {
				fmt.Printf("\n%s=%s\n", p.alias, p.original)
				return alias, original, nil
			}
		}
		return alias, original, fmt.Errorf("%w: alias not found", ErrInvalidArg)

	} else if args[0] == "-r" { // Remove alias
		fmt.Print("\nDeleting alias...\n")
		for i, r := range aliasList {
			if r.alias == args[1] {
				aliasList = append(aliasList[:i], aliasList[i+1:]...)
				alias = append(alias[:i], alias[i+1:]...)
				original = append(original[:i], original[i+1:]...)

				fmt.Printf("Alias deleted.\n")
				return alias, original, nil
			}
		}
		return alias, original, fmt.Errorf("%w: alias not found", ErrInvalidArg)

	}

	// Clean up arguments and append new alias
	// Format: alias [option] [name]='[value]'
	// Example: alias la='ls -a'
	if strings.Contains(args[0], "='") && strings.HasSuffix(args[argsSize-1], "'") {

		// Remove =' from first argument by splitting it
		var split = strings.Split(args[0], "='")

		if len(split) == 2 { // Just two from before and after ='

			// Remove last single quote
			args[argsSize-1] = strings.TrimSuffix(args[argsSize-1], "'")

			// Combine arguments for full original command
			var result []string
			result = append(result, split[0])    // adding alias command
			result = append(result, split[1])    // adding original command
			result = append(result, args[1:]...) // adding rest of original command

			// Add alias pair to list
			aliasList = append(aliasList, AliasPair{alias: result[0], original: result[1:]})
			aliasSize = len(aliasList) //  new alias list size
			alias = append(alias, aliasList[aliasSize-1].alias)
			original = append(original, aliasList[aliasSize-1].original)

			fmt.Printf("\nAdded alias: %s, original: %s\n", aliasList[aliasSize-1].alias, aliasList[aliasSize-1].original)
			return alias, original, nil

		}

	} else {
		fmt.Print(aliasFormat, aliasOptions)
		return alias, original, fmt.Errorf("%w: improper format", ErrInvalidArg)

	}

	fmt.Print(aliasFormat, aliasOptions)
	return alias, original, fmt.Errorf("%w: logical error", ErrInvalidArg)
}
