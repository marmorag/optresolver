package main

import (
	"fmt"
	"github.com/marmorag/optresolver/pkg/optresolver"
	"os"
)

func main() {
	or := optresolver.OptionResolver{
		Options: nil,
		Help:    "This is a program to test option resolver",
	}

	optresolver.AddOption(&or, optresolver.Option{
		Short:    "n",
		Long:     "name",
		Required: false,
		Type:optresolver.ValueType,
		Help:     "A name to display",
	})

	optresolver.AddOption(&or, optresolver.Option{
		Short:    "t",
		Long:     "test",
		Required: false,
		Type:optresolver.ValueType,
		Help:     "A test option",
	})

	optresolver.AddOption(&or, optresolver.Option{
		Short:    "r",
		Long:     "required",
		Required: true,
		Type:optresolver.ValueType,
		Help:     "A required test option",
	})

	optresolver.AddOption(&or, optresolver.Option{
		Short:    "z",
		Long:     "zest",
		Required: false,
		Type:optresolver.BoolType,
		Help:     "Fo further implementation",
	})

	opt, err := optresolver.Parse(or, os.Args)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s : %s", or.Help, err))
		os.Exit(1)
	}

	fmt.Println(opt)
}
