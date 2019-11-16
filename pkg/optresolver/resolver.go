package optresolver

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func (or *OptionResolver) AddOption(opt Option) {
	if opt.Short == "h" || opt.Long == "help" {
		fmt.Println(fmt.Errorf("argument h or help is reserved"))
		os.Exit(1)
	}

	or.Options = append(or.Options, opt)

	if opt.Required {
		or.requiredOptions = append(or.requiredOptions, &opt)
	}

	if opt.Default != "" {
		or.defaultedOptions = append(or.defaultedOptions, &opt)
	}
}

func (or *OptionResolver) Parse(args []string) (map[string]string, error) {
	var currentOption Option
	var isKnownOption bool
	result := make(map[string]string)

	for i, arg := range args {
		if i == 0 {
			continue
		}

		// TODO : handle bool type
		if i % 2 != 0 {
			if arg == "-h" || arg == "--help" {
				or.Help()
				os.Exit(0)
			}

			currentOption, isKnownOption = or.getOpt(arg)

			if !isKnownOption {
				return map[string]string{}, errors.New(fmt.Sprintf("error : unknown option : %s", arg))
			}
		} else if isKnownOption {
			if currentOption.Type == ValueType {
				result[currentOption.Long] = arg
			}
		}
	}

	if requiredOptions, hasRequired := or.hasReqOpts(); hasRequired {
		for _, reqOpt := range requiredOptions {
			if _, exist := result[reqOpt.Long]; !exist && reqOpt.Type != BoolType {
				return map[string]string{}, errors.New(fmt.Sprintf("The flag : %s is required", reqOpt.Long))
			}
		}
	}

	if defaultedOptions, hasDefaults := or.hasDefOpts(); hasDefaults {
		for _, defOpt := range defaultedOptions {
			if _, exist := result[defOpt.Long]; !exist {
				result[defOpt.Long] = defOpt.Default
			}
		}
	}

	return result, nil
}

func (or *OptionResolver) Help() {
	fmt.Println(or.Description)
	fmt.Println("")
	fmt.Println(strings.Repeat("=", len(or.Description)))
	for _, option := range or.Options  {
		fmt.Println(fmt.Sprintf("-%-5s, --%-15s | %s", option.Short, option.Long, option.Help))
		fmt.Println("")
	}

	fmt.Println(fmt.Sprintf("%-6s, %-17s | %s", "-h", "--help", "Display help"))
}