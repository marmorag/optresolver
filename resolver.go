package optresolver

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const ErrorReservedArgument string = "argument h or help is reserved"
const ErrorExistingOption string = "option %s is already registered"
const ErrorMissingOption string = "the flag : %s is required"
const ErrorUnknownOption = "unknown option : %s"

func (or *OptionResolver) AddOption(opt Option) error{
	if opt.Short == "h" || opt.Long == "help" {
		return errors.New(ErrorReservedArgument)
	}

	_, existOption := or.getOpt(opt.Short)

	if existOption {
		return errors.New(fmt.Sprintf(ErrorExistingOption, opt.Long))
	}

	or.Options = append(or.Options, opt)

	if opt.Required {
		or.requiredOptions = append(or.requiredOptions, &opt)
	}

	if opt.Default != "" {
		or.defaultedOptions = append(or.defaultedOptions, &opt)
	}

	return nil
}

func (or *OptionResolver) Resolve(args []string) (map[string]string, error) {
	var currentOption Option
	var isKnownOption bool
	result := make(map[string]string)

	for i, arg := range args {
		if i == 0 {
			continue
		}

		if arg == "-h" || arg == "--help" {
			or.Help()
			os.Exit(0)
		}

		// TODO : handle bool type
		if i % 2 != 0 {
			currentOption, isKnownOption = or.getOpt(arg)

			if !isKnownOption {
				return map[string]string{}, errors.New(fmt.Sprintf(ErrorUnknownOption, arg))
			}
		} else if isKnownOption {
			if currentOption.Type == ValueType {
				result[currentOption.Long] = arg
			}

			isKnownOption = false
		}
	}

	if requiredOptions, hasRequired := or.hasRequiredOptions(); hasRequired {
		for _, reqOpt := range requiredOptions {
			if _, exist := result[reqOpt.Long]; !exist && reqOpt.Type != BoolType {
				return map[string]string{}, errors.New(fmt.Sprintf(ErrorMissingOption, reqOpt.Long))
			}
		}
	}

	if defaultedOptions, hasDefaults := or.hasDefaultOptions(); hasDefaults {
		for _, defOpt := range defaultedOptions {
			if _, exist := result[defOpt.Long]; !exist {
				result[defOpt.Long] = defOpt.Default
			}
		}
	}

	return result, nil
}

func (or *OptionResolver) Help() {
	fmt.Printf("%s\n\n", or.Description)
	fmt.Println(strings.Repeat("=", len(or.Description)))
	for _, option := range or.Options  {
		var reqString string
		var defString string

		if option.Required {
			reqString = fmt.Sprintf("|required")
		} else {
			reqString = ""
		}

		if option.Default != "" {
			defString = fmt.Sprintf("|default : %s", option.Default)
		} else {
			defString = ""
		}

		fmt.Println(fmt.Sprintf("-%-5s, --%-15s %s%s| %s\n", option.Short, option.Long, reqString, defString, option.Help))
	}

	fmt.Println(fmt.Sprintf("%-6s, %-17s | %s", "-h", "--help", "Display help"))
}