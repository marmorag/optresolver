package optresolver

import "strings"

func getOpt(value string, resolver OptionResolver) (Option, bool) {
	for _, opt := range resolver.Options {
		cleaned := strings.Replace(value, "-", "", -1)

		if cleaned == opt.Short || cleaned == opt.Long {
			return opt, true
		}
	}

	return Option{}, false
}

func (or *OptionResolver) hasReqOpts() ([]*Option, bool) {
	if len(or.requiredOptions) > 0 {
		return or.requiredOptions, true
	}

	return []*Option{}, false
}

func (or *OptionResolver) hasDefOpts() ([]*Option, bool) {
	if len(or.defaultedOptions) > 0 {
		return or.defaultedOptions, true
	}

	return []*Option{}, false
}
