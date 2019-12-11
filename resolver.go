package optresolver

import (
	"errors"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
	"strings"
)

const ErrorReservedArgument string = "argument h or help is reserved"
const ErrorExistingOption string = "option %s is already registered"
const ErrorMissingOption string = "the flag : %s is required"
const ErrorUnknownOption string = "unknown option : %s"

type OptionType int

const (
	ValueType OptionType = iota
	BoolType  OptionType = iota
)

type Option struct {
	Short    string
	Long     string
	Required bool
	Type     OptionType
	Default  interface{}
	Help     string
}

type OptionResolver struct {
	Options     []Option
	Description string
	Name        string

	generateAscii    bool
	requiredOptions  []*Option
	defaultedOptions []*Option
	asciiFont        string
}

func NewOptionResolver(name string, description string) *OptionResolver {
	return &OptionResolver{
		Options:          make([]Option, 0),
		Name:             name,
		Description:      description,
		generateAscii:    false,
		requiredOptions:  make([]*Option, 0),
		defaultedOptions: make([]*Option, 0),
	}
}

func (or *OptionResolver) EnableAsciiArt()  {
	or.generateAscii = true
	or.asciiFont = "cybermedium"
}

func (or *OptionResolver) SetAsciiArtFont(font string)  {
	if font != "" {
		or.asciiFont = font
	}
}

func (or *OptionResolver) AddOption(opt Option) error {
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

	if opt.Default != nil {
		or.defaultedOptions = append(or.defaultedOptions, &opt)
	}

	return nil
}

func (or *OptionResolver) Resolve(args []string) (map[string]interface{}, error) {
	var currentOption Option
	var isKnownOption bool
	result := make(map[string]interface{})
	skipArg := false

	mappedArgs := args[1:]

	for i, arg := range mappedArgs {
		if arg == "-h" || arg == "--help" {
			or.Help()
			os.Exit(0)
		}

		if !skipArg {
			currentOption, isKnownOption = or.getOpt(arg)

			if isKnownOption {
				if currentOption.Type == ValueType {
					result[currentOption.Long] = mappedArgs[i+1]
					skipArg = true
				} else if currentOption.Type == BoolType {
					result[currentOption.Long] = true
				}
			} else {
				return make(map[string]interface{}), errors.New(fmt.Sprintf(ErrorUnknownOption, arg))
			}
		} else {
			skipArg = false
		}
	}

	if requiredOptions, hasRequired := or.hasRequiredOptions(); hasRequired {
		for _, reqOpt := range requiredOptions {
			if _, exist := result[reqOpt.Long]; !exist {
				return make(map[string]interface{}), errors.New(fmt.Sprintf(ErrorMissingOption, reqOpt.Long))
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
	var name string
	if or.generateAscii == true {
		asciiArt := figure.NewFigure(or.Name, or.asciiFont, true)
		name = asciiArt.String()
	} else {
		name = or.Name
	}

	fmt.Printf("%s\n", name)
	fmt.Printf("%s\n\n", or.Description)
	fmt.Println(strings.Repeat("=", len(or.Description)))
	for _, option := range or.Options {
		var reqString string
		var defString string

		if option.Required {
			reqString = fmt.Sprintf("|required")
		} else {
			reqString = ""
		}

		if option.Default != nil {
			if option.Type == ValueType {
				defString = fmt.Sprintf("|default : %s", option.Default)
			} else if option.Type == BoolType {
				defString = fmt.Sprintf("|default : %v", option.Default)
			}
		} else {
			defString = ""
		}

		fmt.Println(fmt.Sprintf("-%-5s, --%-15s %s%s| %s\n", option.Short, option.Long, reqString, defString, option.Help))
	}

	fmt.Println(fmt.Sprintf("%-6s, %-17s | %s", "-h", "--help", "Display help"))
}

func (or *OptionResolver) getOpt(value string) (Option, bool) {
	cleaned := strings.Replace(value, "-", "", -1)

	for _, opt := range or.Options {
		if cleaned == opt.Short || cleaned == opt.Long {
			return opt, true
		}
	}

	return Option{}, false
}

func (or *OptionResolver) hasRequiredOptions() ([]*Option, bool) {
	if len(or.requiredOptions) > 0 {
		return or.requiredOptions, true
	}

	return []*Option{}, false
}

func (or *OptionResolver) hasDefaultOptions() ([]*Option, bool) {
	if len(or.defaultedOptions) > 0 {
		return or.defaultedOptions, true
	}

	return []*Option{}, false
}
