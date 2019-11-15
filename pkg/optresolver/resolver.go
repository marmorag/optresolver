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
}

func (or *OptionResolver) Parse(args []string) (map[string]string, error) {
	var currOpt Option
	var isOpt bool
	res := make(map[string]string)

	for i, value := range args {
		if i == 0 {
			continue
		}

		// TODO : handle bool type
		if i % 2 != 0 {
			if value == "-h" || value == "--help" {
				or.Help()
				os.Exit(0)
			}

			currOpt, isOpt = getOpt(value, *or)

			if !isOpt {
				return map[string]string{}, errors.New(fmt.Sprintf("error : unknown option : %s", value))
			}
		} else if isOpt {
			if currOpt.Type == ValueType {
				res[currOpt.Long] = value
			}
		}
	}

	if reqOpts, req := hasReqOpts(*or); req {
		for _, reqOpt := range reqOpts {
			if _, exist := res[reqOpt.Long]; !exist && reqOpt.Type != BoolType {
				return map[string]string{}, errors.New(fmt.Sprintf("The flag : %s is required", reqOpt.Long))
			}
		}
	}

	if defOpts, def := hasDefOpts(*or); def {
		for _, defOpt := range defOpts {
			if _, exist := res[defOpt.Long]; !exist {
				res[defOpt.Long] = defOpt.Default
			}
		}
	}

	return res, nil
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