package optresolver

import (
	"fmt"
	"os"
	"testing"
)

func TestOptionResolver_AddOption(t *testing.T) {

}

func main() {
	or := &OptionResolver{
		Options:     nil,
		Description: "This is a program to test option resolver",
	}

	or.AddOption(Option{
		Short:    "n",
		Long:     "name",
		Required: false,
		Type:     ValueType,
		Default:  "",
		Help:     "A name to display",
	})

	or.AddOption(Option{
		Short:    "t",
		Long:     "test",
		Required: false,
		Type:     ValueType,
		Default:  "default_value",
		Help:     "A test option",
	})

	or.AddOption(Option{
		Short:    "r",
		Long:     "required",
		Required: true,
		Type:     ValueType,
		Help:     "A required test option",
	})

	or.AddOption(Option{
		Short:    "d",
		Long:     "default",
		Required: false,
		Default: "default_value",
		Help:     "A default value",
	})

	opt, err := or.Resolve(os.Args)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s : %s", or.Description, err))
		os.Exit(1)
	}

	fmt.Println(opt)
}