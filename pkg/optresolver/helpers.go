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

func hasReqOpts(resolver OptionResolver) (options []Option, hasRequired bool) {
	for _, opt := range resolver.Options{
		if opt.Required {
			hasRequired = true
			options = append(options, opt)
		}
	}
	return
}

func hasDefOpts(resolver OptionResolver) (options []Option, hasDefaults bool) {
	for _, opt := range resolver.Options{
		if opt.Default != "" {
			hasDefaults = true
			options = append(options, opt)
		}
	}
	return
}
