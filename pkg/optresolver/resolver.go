package optresolver

import (
	"errors"
	"fmt"
)

func (or *OptionResolver) AddOption(opt Option) {
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
			currOpt, isOpt = getOpt(value, *or)

			if !isOpt {
				return map[string]string{}, errors.New(fmt.Sprintf("Invalid input value : %s", value))
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

	return res, nil
}
