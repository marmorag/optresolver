package main

import (
	"fmt"
	"github.com/marmorag/optresolver/pkg/optresolver"
	"os"
)

func main() {
	or := &optresolver.OptionResolver{
		Options:     nil,
		Description: "This is a program to test option resolver",
	}

	or.AddOption(optresolver.Option{
		Short:    "n",
		Long:     "name",
		Required: false,
		Type:optresolver.ValueType,
		Default: "",
		Help:     "A name to display",
	})

	or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Required: false,
		Type:optresolver.ValueType,
		Default: "default_value",
		Help:     "A test option",
	})

	or.AddOption(optresolver.Option{
		Short:    "r",
		Long:     "required",
		Required: true,
		Type:optresolver.ValueType,
		Help:     "A required test option",
	})

	or.AddOption(optresolver.Option{
		Short:    "d",
		Long:     "default",
		Required: false,
		Default: "default_value",
		Help:     "A default value",
	})

	opt, err := or.Parse(os.Args)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s : %s", or.Description, err))
		os.Exit(1)
	}

	fmt.Println(opt)
}
