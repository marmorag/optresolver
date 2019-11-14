package optresolver

import "strings"

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
