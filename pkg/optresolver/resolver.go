package optresolver

import (
	"errors"
	"fmt"
	"strings"
)

type OptionType int

const (
	ValueType OptionType = iota
	BoolType OptionType = iota
)


type Option struct {
	Short string
	Long string
	Required bool
	Type OptionType
	Help string
}

type OptionResolver struct {
	Options []Option
	Help string
}

func AddOption(or *OptionResolver, opt Option) {
	or.Options = append(or.Options, opt)
}

func Parse(or OptionResolver, args []string) (map[string]string, error) {
	var currOpt Option
	var isOpt bool
	res := make(map[string]string)

	for i, value := range args {
		if i == 0 {
			continue
		}

		// TODO : handle bool type
		if i % 2 != 0 {
			currOpt, isOpt = getOpt(value, or)

			if !isOpt {
				return map[string]string{}, errors.New(fmt.Sprintf("Invalid input value : %s", value))
			}
		} else if isOpt {
			if currOpt.Type == ValueType {
				res[currOpt.Long] = value
			}
		}
	}

	if reqOpts, req := hasReqOpts(or); req {
		for _, reqOpt := range reqOpts {
			if _, exist := res[reqOpt.Long]; !exist && reqOpt.Type != BoolType {
				return map[string]string{}, errors.New(fmt.Sprintf("The flag : %s is required", reqOpt.Long))
			}
		}
	}

	return res, nil
}

func getOpt(value string, or OptionResolver) (Option, bool) {
	for _, opt := range or.Options {
		cleaned := strings.Replace(value, "-", "", -1)

		if cleaned == opt.Short || cleaned == opt.Long {
			return opt, true
		}
	}

	return Option{}, false
}

func hasReqOpts(or OptionResolver) (opts []Option, hasReq bool) {
	for _, opt := range or.Options{
		if opt.Required {
			hasReq = true
			opts = append(opts, opt)
		}
	}
	return
}
