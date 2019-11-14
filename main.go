package main

import (
	"fmt"
	"github.com/marmorag/optresolver/pkg/optresolver"
	"os"
)

func main() {
	or := &optresolver.OptionResolver{
		Options: nil,
		Help:    "This is a program to test option resolver",
	}

	or.AddOption(optresolver.Option{
		Short:    "n",
		Long:     "name",
		Required: false,
		Type:optresolver.ValueType,
		Help:     "A name to display",
	})

	or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Required: false,
		Type:optresolver.ValueType,
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
		Short:    "z",
		Long:     "zest",
		Required: false,
		Type:optresolver.BoolType,
		Help:     "Fo further implementation",
	})

	opt, err := or.Parse(os.Args)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s : %s", or.Help, err))
		os.Exit(1)
	}

	fmt.Println(opt)
}
