package optresolver

import "strings"

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
